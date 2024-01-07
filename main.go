package main

import (
	"context"
	"database/sql"
	"desktop-cycle/internal/db"
	"embed"

	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed all:migrations
var migrations embed.FS

// create db and run migrations
func setup() (con *sql.DB, err error) {
	err = db.CreateDBFile()
	if err != nil {
		return nil, err
	}

	con, err = db.ConnectToDB()
	if err != nil {
		return nil, err
	}

	fs, err := iofs.New(migrations, "migrations")
	if err != nil {
		return nil, err
	}

	err = db.RunMigrations(con, fs)
	if err != nil {
		return nil, err
	}

	return con, nil
}

func main() {
	con, err := setup()
	if err != nil {
		panic(err)
	}

	defer con.Close()

	app := NewApp()
	cycle := NewCycle()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Desktop Cycle",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			cycle.startup(ctx)
		},
		Bind: []interface{}{
			app,
			cycle,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
