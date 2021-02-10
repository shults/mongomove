package application

import "github.com/shults/mongomove/application/config"

type App struct {
	cnf  *config.Config
	opts *applicationOptions
}

func (a *App) Run() error {
	// todo: add something for system communication and testing
	return nil
}

func NewApp(cnf *config.Config, opts *applicationOptions) *App {
	return &App{cnf, opts}
}
