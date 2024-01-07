package driver

import (
	"desktop-cycle/internal/util"
	"errors"
	"sync"
	"time"

	"github.com/simpleiot/simpleiot/respreader"
	"go.bug.st/serial"
)

const DESK_CYCLE_DRIVER = "DeskCycle"
const deskCycleName = "DeskCycle Speedo\r\n"

var handshakeMsg = []byte("h\r\n")
var speedMsg = []byte("s\r\n")

var mode = &serial.Mode{
	BaudRate: 9600,
	DataBits: 8,
	StopBits: serial.OneStopBit,
	Parity:   serial.NoParity,
}

// DeskCycle wrapper for a ReadWriteCloser port with an open desk cycle serial device.
// Provides methods for getting all desk cycle messages.
type DeskCycle struct {
	port *respreader.ReadWriteCloser
	mu   sync.Mutex
}

// Handshake returns the handshake response from the desk cycle
func (dc *DeskCycle) Handshake() (string, error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return util.ReadLine(dc.port, handshakeMsg, 3)
}

// GetSpeed returns the current speed of the desk cycle
func (dc *DeskCycle) GetSpeed() (float64, error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return util.ReadFloat(dc.port, speedMsg, 3)
}

// Close closes the desk cycle serial device
func (dc *DeskCycle) Close() error {
	return dc.port.Close()
}

// NewDeskCycleSerial search serial ports for a desk cycle.
// Return DeskCycleSerial with the opened port
func NewDeskCycle() (dc *DeskCycle, err error) {
	devNames, err := serial.GetPortsList()

	if err != nil {
		return nil, err
	}

	if len(devNames) == 0 {
		return nil, errors.New("no serial ports found")
	}

	for _, dev := range devNames {
		port, err := serial.Open(dev, mode)
		if err != nil {
			continue
		}

		portTimeout := respreader.NewReadWriteCloser(port, time.Second, time.Millisecond*50)

		response, err := util.ReadLine(portTimeout, handshakeMsg, 3)
		if err != nil {
			continue
		}

		if response == deskCycleName {
			dc = &DeskCycle{port: portTimeout}
			break
		} else {
			err = portTimeout.Close()

			if err != nil {
				return nil, err
			}
		}
	}

	if dc == nil {
		return nil, errors.New("failed to find a DeskCycleSerial device")
	}

	return dc, nil
}
