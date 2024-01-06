package util

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func getDBPath() string {
	// TODO: add cross platform paths
	home, _ := os.UserHomeDir()
	if runtime.GOOS == "linux" {
		return home + "/.local/share/desktop-cycle/db.sqlite3"
	}

	panic("Unsupported OS " + runtime.GOOS)
}

func ConnectToDB() (*sql.DB, error) {
	dbPath := getDBPath()
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB at "+dbPath)
	}

	return db, nil
}

func RunMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("creating sqlite3 db driver failed %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:////home/ethan/git/ethanmassie/desktop-cycle/migrations", "sqlite3", driver)
	if err != nil {
		return fmt.Errorf("initializing db migration failed %s", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrating database failed %s", err)
	}

	return nil
}
