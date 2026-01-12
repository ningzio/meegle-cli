package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/meegle"
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
