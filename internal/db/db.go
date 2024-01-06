package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

const DB_FILE = "db.sqlite3"
const LINUX_PATH = "/.local/share/desktop-cycle/"

func GetDBPath() string {
	// TODO: add cross platform paths
	home, _ := os.UserHomeDir()
	if runtime.GOOS == "linux" {
		return home + LINUX_PATH
	}

	panic("Unsupported OS " + runtime.GOOS)
}

func CreateLocalDirectory() error {
	localPath := GetDBPath()
	err := os.MkdirAll(localPath, os.ModePerm)
	if err != nil {
		return err
	}

	dbFile, err := os.OpenFile(filepath.Join(localPath, DB_FILE), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	return dbFile.Close()
}

func ConnectToDB() (*sql.DB, error) {
	dbPath := GetDBPath()
	db, err := sql.Open("sqlite3", filepath.Join(dbPath, DB_FILE))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB at "+dbPath)
	}

	return db, nil
}

func RunMigrations(db *sql.DB, migrations source.Driver) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("creating sqlite3 db driver failed %s", err)
	}

	m, err := migrate.NewWithInstance("iofs", migrations, "sqlite3", driver)
	if err != nil {
		return fmt.Errorf("initializing db migration failed %s", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrating database failed %s", err)
	}

	return nil
}
