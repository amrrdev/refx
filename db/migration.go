package db

import (
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	migrationsPath := strings.ReplaceAll(wd, "\\", "/") + "/db/migrations"
	sourceURL := "file://" + migrationsPath

	log.Println("Migration path:", sourceURL)

	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("✅ Database migrated")

	return nil
}
