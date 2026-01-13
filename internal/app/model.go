package app

import (
	"sync/atomic"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/meegle"
	"meegle-cli/internal/screen"
	"meegle-cli/internal/screens/tasks"
	"meegle-cli/internal/store"
)

type Config struct {
	ProjectKey string
}

type App struct {
	Router   *Router
	Store    store.State
	Overlays *Overlays
	Cmds     *meegle.Cmds
	KeyMap   KeyMap
	Theme    Theme
	Config   Config
	reqID    int64
}

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

func (a *App) Init() tea.Cmd {
	return a.Router.Current().Init(a)
}

func (a *App) StoreState() store.State {
	return a.Store
}

func (a *App) MeegleCmds() *meegle.Cmds {
	return a.Cmds
}

func (a *App) ProjectKey() string {
	return a.Config.ProjectKey
}

func (a *App) NextReqID() int64 {
	return atomic.AddInt64(&a.reqID, 1)
}

func (a *App) Push(next screen.Screen) tea.Cmd {
	return a.Router.Push(a, next)
}

func (a *App) Pop() tea.Cmd {
	return a.Router.Pop(a)
}

func (a *App) Replace(next screen.Screen) tea.Cmd {
	return a.Router.Replace(a, next)
}
