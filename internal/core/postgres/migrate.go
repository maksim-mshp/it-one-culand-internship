package postgres

import (
	"culand-internship/internal/core/config"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func RunMigrations(config config.Database) error {
	dsn := MakeConnectionString(config)
	m, err := migrate.New(
		"file://internal/core/postgres/migrations",
		dsn,
	)
	if err != nil {
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
