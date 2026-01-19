package core

import (
	"context"
	"culand-internship/internal/core/config"
	"culand-internship/internal/core/database"
	corehttp "culand-internship/internal/core/http"
	internshipApp "culand-internship/internal/internship/app"
	internshipHttp "culand-internship/internal/internship/infra/http"
	internshipPostgres "culand-internship/internal/internship/infra/postgres"
	"database/sql"
	"net/http"
	"time"
)

type App struct {
	Config   *config.Config
	Database *sql.DB
	Server   *http.Server
}

func Start(cfg *config.Config) (*App, error) {
	db, err := database.NewPostgres(cfg.Database)
	if err != nil {
		return nil, err
	}
	if err := database.RunMigrations(db); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	internshipRepo := internshipPostgres.NewRepository(db)
	internshipHandlers := internshipApp.BuildHandlers(internshipRepo)
	internshipHttpHandler := internshipHttp.NewHttpHandler(internshipHandlers)
	internshipHttp.RegisterRoutes(mux, internshipHttpHandler)

	srv, err := corehttp.NewServer(cfg.Port, mux)
	if err != nil {
		return nil, err
	}

	return &App{
		Config:   cfg,
		Database: db,
		Server:   srv,
	}, nil
}

func (a *App) Stop(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	if err := a.Database.Close(); err != nil {
		return err
	}

	return nil
}
