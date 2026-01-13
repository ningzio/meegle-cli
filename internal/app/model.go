package app

import (
	"sync/atomic"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/meegle"
	"meegle-cli/internal/screen"
	"meegle-cli/internal/screens/tasks"
	"meegle-cli/internal/store"
)

// Config captures runtime settings for the app.
type Config struct {
	ProjectKey string
}

// App is the root Bubble Tea model for the Meegle TUI.
type App struct {
	Router   *Router
	Store    store.State
	Overlays *Overlays
	Cmds     *meegle.Cmds
	KeyMap   KeyMap
	Theme    Theme
	Config   Config
	reqID    int64
	width    int
	height   int
}

// New builds a new app model with default dependencies.
func New(config Config, cmds *meegle.Cmds) *App {
	tasksScreen := tasks.New()
	router := NewRouter(tasksScreen)

	return &App{
		Router:   router,
		Store:    store.NewState(),
		Overlays: NewOverlays(),
		Cmds:     cmds,
		KeyMap:   DefaultKeyMap(),
		Theme:    DefaultTheme(),
		Config:   config,
	}
}

// Init initializes the app by delegating to the current screen.
func (a *App) Init() tea.Cmd {
	return a.Router.Current().Init(a)
}

// StoreState returns the current store snapshot.
func (a *App) StoreState() store.State {
	return a.Store
}

// MeegleCmds exposes the Meegle command factory for screens.
func (a *App) MeegleCmds() *meegle.Cmds {
	return a.Cmds
}

// ProjectKey returns the configured project key.
func (a *App) ProjectKey() string {
	return a.Config.ProjectKey
}

// NextReqID returns a new request identifier for async workflows.
func (a *App) NextReqID() int64 {
	return atomic.AddInt64(&a.reqID, 1)
}

// Push navigates to the next screen.
func (a *App) Push(next screen.Screen) tea.Cmd {
	return a.Router.Push(a, next)
}

// Pop navigates back to the previous screen.
func (a *App) Pop() tea.Cmd {
	return a.Router.Pop(a)
}

// Replace swaps the current screen with another.
func (a *App) Replace(next screen.Screen) tea.Cmd {
	return a.Router.Replace(a, next)
}

// WindowSize returns the current window dimensions.
func (a *App) WindowSize() (int, int) {
	return a.width, a.height
}
