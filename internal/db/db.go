package db

import (
	"database/sql"
	"desktop-cycle/internal/util"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

const DB_FILE = "db.sqlite3"

// create the sqlite file if it doesn't exist
func CreateDBFile() error {
	localPath, err := util.GetLocalPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(localPath, os.ModePerm)
	if err != nil {
		return err
	}

	dbFile, err := os.OpenFile(filepath.Join(localPath, DB_FILE), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	return dbFile.Close()
}

// create connection to the sqlite db
func ConnectToDB() (*sql.DB, error) {
	dbPath, err := util.GetLocalPath()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", filepath.Join(dbPath, DB_FILE))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite DB at "+dbPath)
	}

	return db, nil
}

// run the given migrations on the given sqlite3 db
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
