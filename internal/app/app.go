package app

import (
	"context"

	"github.com/mixdjoker/server-template/internal/closer"
	"github.com/mixdjoker/server-template/internal/config"
	"github.com/mixdjoker/server-template/internal/logger"
	"github.com/spf13/pflag"
)

// App is a struct that holds all the application dependencies
type App struct {
	serviceProvider *serviceProvider
	// Servers
}

// NewApp is a function that returns a new instance of the App struct
func NewApp(ctx context.Context) (*App, error) {
	a := App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		// Init functions
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	envFilePath := pflag.StringP("env", "e", ".env", ".env file path")
	pflag.Parse()

	if err := config.Load(*envFilePath); err != nil {
		return err
	}

	return nil
}

func (a *App) initLogger(ctx context.Context) error {
	cfg, err := config.NewLoggerConfig()
	if err != nil {
		return err
	}

	logger.SetupLoggerLevel(cfg.Level(), cfg.Format())
	logger.Info(ctx).Msg("Logger initialized")

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// Run starts initialised application.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return nil
}
