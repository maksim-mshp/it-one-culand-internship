package database

import (
	"culand-internship/internal/core/config"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgres(config config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		if err := db.Close(); err != nil {
			return nil, fmt.Errorf("failed to close db: %w", err)
		}
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}
	return db, nil
}
