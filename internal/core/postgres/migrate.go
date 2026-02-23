package postgres

import (
	"culand-internship/internal/core/config"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func RunMigrations(config config.Database) error {
	dsn := MakeConnectionString(config)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		_ = db.Close()
		return fmt.Errorf("failed to ping database: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		_ = db.Close()
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/core/postgres/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		_ = db.Close()
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	defer func() {
		sourceErr, dbErr := m.Close()
		if sourceErr != nil {
			log.Printf("failed close migration source: %v", sourceErr)
		}
		if dbErr != nil {
			log.Printf("failed close migration database: %v", dbErr)
		}
	}()

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("migrations: no changes detected")
			return nil
		}
		return fmt.Errorf("migrations failed: %w", err)
	}

	log.Println("migrations applied successfully")
	return nil
}
