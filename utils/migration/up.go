package migration

import (
	"errors"
	"fmt"

	"github.com/JesseNicholas00/GogoManager/utils/logging"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp(dbString string, migrationsPath string) error {
	migrationLogger := logging.GetLogger("migration", "up")

	sourceURL := fmt.Sprintf("file://%s", migrationsPath)
	m, err := migrate.New(sourceURL, dbString)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		migrationLogger.Info("no migration changes")
	} else {
		migrationLogger.Info("database migrated")
	}

	return nil
}
