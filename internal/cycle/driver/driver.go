package driver

import "errors"

type CycleDriver interface {
	GetSpeed() (float64, error)
	Close() error
}

func GetDriver(name string) (CycleDriver, error) {
	if name == DESK_CYCLE_DRIVER {
		return NewDeskCycle()
	}

	return nil, errors.New("driver not found")
}
