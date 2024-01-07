package main

import (
	"context"
	"desktop-cycle/internal/cycle/driver"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Cycle struct
type Cycle struct {
	ctx           context.Context
	cycleDriver   driver.CycleDriver
	stopSpeedPoll chan struct{}
}

// NewApp creates a new App application struct
func NewCycle() *Cycle {
	return &Cycle{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
//
//lint:ignore U1000 called by framework
func (c *Cycle) startup(ctx context.Context) {
	c.ctx = ctx
}

//lint:ignore U1000 called by framework
func (c *Cycle) beforeClose(ctx context.Context) (prevent bool) {
	if c.cycleDriver != nil {
		c.stopSpeedPoll <- struct{}{}
		c.cycleDriver.Close()
	}

	return false
}

func (c *Cycle) Connect() int {
	if c.cycleDriver != nil {
		return 2
	}

	dc, err := driver.GetDriver(driver.DESK_CYCLE_DRIVER)

	if err != nil {
		fmt.Println(err)
		return 1
	}

	c.cycleDriver = dc
	c.startSpeedPoll()

	return 0
}

func (c *Cycle) startSpeedPoll() {
	poll := time.NewTicker(250 * time.Millisecond)
	c.stopSpeedPoll = make(chan struct{})

	go func() {
		for {
			select {
			case <-poll.C:
				speed, err := c.cycleDriver.GetSpeed()
				if err != nil {
					runtime.LogError(c.ctx, fmt.Sprintf("Error getting speed: %s", err))
				}

				runtime.EventsEmit(c.ctx, "speed", speed)
			case <-c.stopSpeedPoll:
				poll.Stop()
				return
			}
		}
	}()
}
