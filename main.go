package main

import (
	"database/sql"
	"desktop-cycle/internal/util"
	"embed"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func createLocalDirectory() error {
	// TODO: add cross platform paths
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	localPath := filepath.Join(home, ".local", "share", "desktop-cycle")
	err = os.MkdirAll(localPath, os.ModePerm)
	if err != nil {
		return err
	}

	db, err := os.Create(filepath.Join(localPath, "db.sqlite3"))
	if err != nil {
		return err
	}

	return db.Close()
}

func setup() (db *sql.DB, err error) {
	err = createLocalDirectory()
	if err != nil {
		return nil, err
	}

	db, err = util.ConnectToDB()
	if err != nil {
		return nil, err
	}

	err = util.RunMigrations(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := setup()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "desktop-cycle",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
