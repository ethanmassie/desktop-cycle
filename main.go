package main

import (
	"desktop-cycle/internal/util"
	"github.com/go-vgo/robotgo"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

var deskCycleSerial *util.DeskCycleSerial

func getSpeed() float64 {
	speed, err := deskCycleSerial.CurrentSpeed()
	if err != nil {
		panic(err)
	}

	return speed
}

func getCadence() float64 {
	cadence, err := deskCycleSerial.CurrentCadence()
	if err != nil {
		panic(err)
	}

	return cadence
}

func toggleKey(keyName string, action string) {
	robotgo.KeyToggle(keyName, action)
}

func tapKey(keyName string) {
	robotgo.KeyTap(keyName)
}

func main() {
	html := mewn.String("./frontend/dist/my-app/index.html")
	js := mewn.String("./frontend/dist/my-app/main.js")
	css := mewn.String("./frontend/dist/my-app/styles.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Desktop Cycle",
		HTML: 	html,
		JS:     js,
		CSS:    css,
		Resizable: true,
	})

	// initialize desk cycle serial
	var err error
	deskCycleSerial, err = util.NewDeskCycleSerial()
	if err != nil {
		panic(err)
	}

	app.Bind(getSpeed)
	app.Bind(getCadence)
	app.Bind(tapKey)
	app.Bind(toggleKey)
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
