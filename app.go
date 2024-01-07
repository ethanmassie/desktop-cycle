package main

import (
	"context"
	"desktop-cycle/internal/util"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx           context.Context
	deskCycle     *util.DeskCycleSerial
	stopSpeedPoll chan struct{}
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

//lint:ignore U1000 called by framework
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	if a.deskCycle != nil {
		a.stopSpeedPoll <- struct{}{}
		a.deskCycle.Close()
	}

	return false
}

func (a *App) Connect() int {
	if a.deskCycle != nil {
		return 2
	}

	dc, err := util.NewDeskCycleSerial()

	if err != nil {
		fmt.Println(err)
		return 1
	}

	a.deskCycle = dc
	a.startSpeedPoll()

	return 0
}

func (a *App) startSpeedPoll() {
	poll := time.NewTicker(250 * time.Millisecond)
	pollStop := make(chan struct{})

	go func() {
		for {
			select {
			case <-poll.C:
				speed, err := a.deskCycle.CurrentSpeed()
				if err != nil {
					fmt.Printf("Error getting speed: %s", err)
				}

				runtime.EventsEmit(a.ctx, "speed", speed)
			case <-pollStop:
				poll.Stop()
				return
			}
		}
	}()

	a.stopSpeedPoll = pollStop
}
