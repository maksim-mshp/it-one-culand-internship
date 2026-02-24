package postgres

import (
	"context"
	"culand-internship/internal/core/config"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/url"
)

func MakeConnectionString(config config.Database) string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.User, config.Password),
		Host:   fmt.Sprintf("%s:%d", config.Host, config.Port),
		Path:   config.Database,
	}
	return u.String()
}

func NewPostgres(config config.Database) (*pgxpool.Pool, error) {
	dsn := MakeConnectionString(config)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}
	return pool, nil
}
