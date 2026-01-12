package app

import tea "github.com/charmbracelet/bubbletea"

type Screen interface {
	Init(app *App) tea.Cmd
	Update(app *App, msg tea.Msg) tea.Cmd
	View(app *App) string

	OnFocus(app *App) tea.Cmd
	OnBlur(app *App)
}

type Router struct {
	stack []Screen
}

func NewRouter(initial Screen) *Router {
	return &Router{stack: []Screen{initial}}
}

func (r *Router) Current() Screen {
	return r.stack[len(r.stack)-1]
}

func (r *Router) Push(app *App, screen Screen) tea.Cmd {
	current := r.Current()
	current.OnBlur(app)
	r.stack = append(r.stack, screen)
	return screen.OnFocus(app)
}

func (r *Router) Pop(app *App) tea.Cmd {
	if len(r.stack) <= 1 {
		return nil
	}

	current := r.Current()
	current.OnBlur(app)
	r.stack = r.stack[:len(r.stack)-1]
	return r.Current().OnFocus(app)
}

func (r *Router) Replace(app *App, screen Screen) tea.Cmd {
	current := r.Current()
	current.OnBlur(app)
	r.stack[len(r.stack)-1] = screen
	return screen.OnFocus(app)
}
