package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"meegle-cli/internal/adapter/lark"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

// ProjectItem implements list.Item interface
type ProjectItem struct {
	id, title, desc string
}

func (i ProjectItem) Title() string       { return i.title }
func (i ProjectItem) Description() string { return i.desc }
func (i ProjectItem) FilterValue() string { return i.title }

type ProjectListModel struct {
	list     list.Model
	projects []lark.Project
	err      error
}

func NewProjectListModel(projects []lark.Project) ProjectListModel {
	items := make([]list.Item, len(projects))
	for i, p := range projects {
		items[i] = ProjectItem{
			id:    p.ID,
			title: p.Name,
			desc:  fmt.Sprintf("Key: %s", p.Key),
		}
	}

	// Setup list
	delegate := list.NewDefaultDelegate()
	l := list.New(items, delegate, 0, 0)
	l.Title = "Select Project"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)

	return ProjectListModel{
		list:     l,
		projects: projects,
	}
}

func (m ProjectListModel) Init() tea.Cmd {
	return nil
}

func (m ProjectListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			// Handle selection
			// i, ok := m.list.SelectedItem().(ProjectItem)
			// if ok {
			// 	m.choice = i
			// }
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ProjectListModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v", m.err)
	}
	return docStyle.Render(m.list.View())
}
