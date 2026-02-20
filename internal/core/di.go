package core

import (
	"context"
	"culand-internship/internal/core/config"
	corehttp "culand-internship/internal/core/http"
	"culand-internship/internal/core/postgres"
	internshipApp "culand-internship/internal/internship/app/handlers"
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
	db, err := postgres.NewPostgres(cfg.Database)
	if err != nil {
		return nil, err
	}
	if err = postgres.RunMigrations(cfg.Database); err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	coreTxManager := postgres.NewPgxTxManager(db)

	adminMW := corehttp.RequireAdminMiddleware()

	internshipRepo := internshipPostgres.NewRepository(db)
	internshipTxRunner := internshipPostgres.NewTxRunner(coreTxManager)
	internshipHandlers := internshipApp.BuildHandlers(internshipRepo, internshipTxRunner)
	internshipV1HttpHandler := internshipV1Http.NewHttpHandler(internshipHandlers)
	internshipV1Http.RegisterRoutes(mux, internshipV1HttpHandler, adminMW)

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
