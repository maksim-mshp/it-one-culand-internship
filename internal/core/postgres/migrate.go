package postgres

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	pgxmigrate "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"log"
)

func RunMigrations(pool *pgxpool.Pool) error {
	sqlDB := stdlib.OpenDBFromPool(pool)
	needCloseSQLDB := true
	defer func() {
		if !needCloseSQLDB {
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("failed close db: %v", err)
		}
	}()

	driver, err := pgxmigrate.WithInstance(sqlDB, &pgxmigrate.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migrate driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/core/postgres/migrations",
		"pgx",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	needCloseSQLDB = false
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
