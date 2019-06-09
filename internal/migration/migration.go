package migration

import (
	"fmt"
	"os"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

// Migrate runs database migrations.
//
// path: path to migrations directory
// args: migration tool arguments
func Migrate(migrationDir string, args ...string) error {
	db := pg.Connect(&pg.Options{
		Database: os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
		Addr:     os.Getenv("DATABASE_URL"),
	})

	// load migration files
	m := migrations.NewCollection()
	m.DisableSQLAutodiscover(false)
	err := m.DiscoverSQLMigrations(migrationDir)
	if err != nil {
		return err
	}

	// start migration
	oldVersion, newVersion, err := m.Run(db, args...)
	if err != nil {
		return err
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("current version is %d\n", oldVersion)
	}
	return nil
}
