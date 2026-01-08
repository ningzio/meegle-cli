package app

import tea "github.com/charmbracelet/bubbletea"

type Screen interface {
	ID() string
	Init() tea.Cmd
	Update(msg tea.Msg, app *Model) tea.Cmd
	View(app *Model) string
}

type Router struct {
	stack []Screen
}

func NewRouter(initial Screen) *Router {
	return &Router{stack: []Screen{initial}}
}

func (r *Router) Current() Screen {
	if len(r.stack) == 0 {
		return nil
	}
	return r.stack[len(r.stack)-1]
}

func (r *Router) Push(screen Screen) {
	r.stack = append(r.stack, screen)
}

func (r *Router) Pop() {
	if len(r.stack) == 0 {
		return
	}
	r.stack = r.stack[:len(r.stack)-1]
}

func (r *Router) GoTo(screen Screen) {
	r.stack = []Screen{screen}
}
