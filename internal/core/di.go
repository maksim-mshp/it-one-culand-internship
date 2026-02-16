package core

import (
	"context"
	"culand-internship/internal/core/config"
	"culand-internship/internal/core/database"
	corehttp "culand-internship/internal/core/http"
	internshipApp "culand-internship/internal/internship/app"
	internshipV1Http "culand-internship/internal/internship/infra/http/v1"
	internshipPostgres "culand-internship/internal/internship/infra/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"time"
)

type App struct {
	Config   *config.Config
	Database *pgxpool.Pool
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
	internshipV1HttpHandler := internshipV1Http.NewHttpHandler(internshipHandlers)
	internshipV1Http.RegisterRoutes(mux, internshipV1HttpHandler)

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

	a.Database.Close()
	return nil
}
