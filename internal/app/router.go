package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
)

type Router struct {
	stack []screen.Screen
}

func NewRouter(initial screen.Screen) *Router {
	return &Router{stack: []screen.Screen{initial}}
}

func (r *Router) Current() screen.Screen {
	return r.stack[len(r.stack)-1]
}

func (r *Router) Push(app screen.AppModel, next screen.Screen) tea.Cmd {
	current := r.Current()
	current.OnBlur(app)
	r.stack = append(r.stack, next)
	return next.OnFocus(app)
}

func (r *Router) Pop(app screen.AppModel) tea.Cmd {
	if len(r.stack) <= 1 {
		return nil
	}

	current := r.Current()
	current.OnBlur(app)
	r.stack = r.stack[:len(r.stack)-1]
	return r.Current().OnFocus(app)
}

func (r *Router) Replace(app screen.AppModel, next screen.Screen) tea.Cmd {
	current := r.Current()
	current.OnBlur(app)
	r.stack[len(r.stack)-1] = next
	return next.OnFocus(app)
}
