package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"meegle-cli/internal/meegle"
	"meegle-cli/internal/store"
	"meegle-cli/internal/ui/components/modal"
	"meegle-cli/internal/ui/components/toast"
)

type Model struct {
	Router *Router
	Store  store.State
	Client meegle.Client
	Theme  Theme
	Toast  toast.Model
	Modal  modal.Model
}

func NewModel(router *Router, client meegle.Client, state store.State) Model {
	return Model{
		Router: router,
		Store:  state,
		Client: client,
		Theme:  DefaultTheme(),
		Toast:  toast.Model{},
		Modal:  modal.Model{},
	}
}

func (m Model) Init() tea.Cmd {
	if m.Router == nil {
		return nil
	}
	if current := m.Router.Current(); current != nil {
		return current.Init()
	}
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	var toastCmd tea.Cmd
	m.Toast, toastCmd = m.Toast.Update(msg)
	if toastCmd != nil {
		cmds = append(cmds, toastCmd)
	}

	var modalCmd tea.Cmd
	if m.Modal.Visible {
		m.Modal, modalCmd = m.Modal.Update(msg)
		if modalCmd != nil {
			cmds = append(cmds, modalCmd)
		}
		if _, ok := msg.(tea.KeyMsg); ok {
			return m, tea.Batch(cmds...)
		}
	}

	switch typed := msg.(type) {
	case ToastMsg:
		var kind toast.Kind
		if typed.Kind == ToastError {
			kind = toast.Error
		}
		m.Toast, toastCmd = m.Toast.Show(typed.Text, kind)
		cmds = append(cmds, toastCmd)
	case ShowConfirmMsg:
		m.Modal = m.Modal.Show(typed.Title, typed.Body, typed.ConfirmText, typed.CancelText, typed.Danger, typed.OnConfirm)
	case PushScreenMsg:
		m.Router.Push(typed.Screen)
		cmds = append(cmds, typed.Screen.Init())
	case PopScreenMsg:
		m.Router.Pop()
	case GoToScreenMsg:
		m.Router.GoTo(typed.Screen)
		cmds = append(cmds, typed.Screen.Init())
	case TriggerFetchTasksMsg:
		reqID := store.NextReqID(&m.Store, store.ReqFetchTasks)
		cmds = append(cmds, meegle.FetchTasksCmd(m.Client, reqID))
	case TriggerCreateTaskMsg:
		reqID := store.NextReqID(&m.Store, store.ReqCreateTask)
		cmds = append(cmds, meegle.CreateTaskCmd(m.Client, reqID, typed.Name))
	case TriggerCreateSubTaskMsg:
		reqID := store.NextReqID(&m.Store, store.ReqCreateSubTask)
		cmds = append(cmds, meegle.CreateSubTaskCmd(m.Client, reqID, typed.TaskID, typed.Name))
	case TriggerToggleSubTaskMsg:
		reqID := store.NextReqID(&m.Store, store.ReqToggleSubTask)
		cmds = append(cmds, meegle.ToggleSubTaskDoneCmd(m.Client, reqID, typed.TaskID, typed.SubTaskID, typed.Done))
	}

	m.Store = store.Reduce(m.Store, msg)

	switch typed := msg.(type) {
	case store.TasksFetchedMsg:
		if typed.Err != nil {
			m.Toast, toastCmd = m.Toast.Show("Failed to fetch tasks: "+typed.Err.Error(), toast.Error)
			cmds = append(cmds, toastCmd)
		}
	case store.TaskCreatedMsg:
		if typed.Err != nil {
			m.Toast, toastCmd = m.Toast.Show("Failed to create task: "+typed.Err.Error(), toast.Error)
			cmds = append(cmds, toastCmd)
		}
	case store.SubTaskCreatedMsg:
		if typed.Err != nil {
			m.Toast, toastCmd = m.Toast.Show("Failed to create subtask: "+typed.Err.Error(), toast.Error)
			cmds = append(cmds, toastCmd)
		}
	case store.SubTaskToggledMsg:
		if typed.Err != nil {
			m.Toast, toastCmd = m.Toast.Show("Failed to update subtask: "+typed.Err.Error(), toast.Error)
			cmds = append(cmds, toastCmd)
		}
	}

	if current := m.Router.Current(); current != nil {
		cmds = append(cmds, current.Update(msg, &m))
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	current := m.Router.Current()
	view := ""
	if current != nil {
		view = current.View(&m)
	}
	if toastView := m.Toast.View(m.toastStyle()); toastView != "" {
		view = view + "\n" + toastView
	}
	if m.Modal.Visible {
		view = lipOverlay(view, m.Modal.View(m.modalTheme()))
	}
	return view
}

func (m Model) toastStyle() lipgloss.Style {
	if m.Toast.Kind == toast.Error {
		return m.Theme.ToastError
	}
	return m.Theme.Toast
}

func (m Model) modalTheme() modal.Theme {
	button := m.Theme.ModalButton
	if m.Modal.Danger {
		button = m.Theme.ModalDanger
	}
	return modal.Theme{
		Frame:  m.Theme.Modal,
		Title:  m.Theme.ModalTitle,
		Body:   m.Theme.Muted,
		Button: button,
		Muted:  m.Theme.Muted,
	}
}

func lipOverlay(base, overlay string) string {
	if overlay == "" {
		return base
	}
	return base + "\n\n" + overlay
}
