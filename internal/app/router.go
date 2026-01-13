package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
)

// Router manages the navigation stack for screens.
type Router struct {
	stack []screen.Screen
}

// NewRouter creates a router with an initial screen on the stack.
func NewRouter(initial screen.Screen) *Router {
	return &Router{stack: []screen.Screen{initial}}
}

// Current returns the top-most screen on the stack.
// It panics if the stack is empty.
func (r *Router) Current() screen.Screen {
	return r.stack[len(r.stack)-1]
}

// Push places a new screen on the stack and focuses it.
func (r *Router) Push(app screen.AppModel, next screen.Screen) tea.Cmd {
	current := r.Current()
	current.OnBlur(app)
	r.stack = append(r.stack, next)
	return next.OnFocus(app)
}

// Pop removes the current screen and focuses the previous one.
func (r *Router) Pop(app screen.AppModel) tea.Cmd {
	if len(r.stack) <= 1 {
		return nil
	}

	current := r.Current()
	current.OnBlur(app)
	r.stack = r.stack[:len(r.stack)-1]
	return r.Current().OnFocus(app)
}

// Replace swaps the current screen while preserving stack depth.
func (r *Router) Replace(app screen.AppModel, next screen.Screen) tea.Cmd {
	current := r.Current()
	current.OnBlur(app)
	r.stack[len(r.stack)-1] = next
	return next.OnFocus(app)
}
