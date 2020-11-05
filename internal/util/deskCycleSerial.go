package util

import (
	"bytes"
	"errors"
	"github.com/simpleiot/simpleiot/respreader"
	"go.bug.st/serial"
	"strconv"
	"strings"
	"sync"
	"time"
)

const deskCycleName = "DeskCycle Speedo\r\n"
const newLineStr = "\r\n"

var handshakeMsg = []byte("h\r\n")
var speedMsg = []byte("s\r\n")
var cadenceMsg = []byte("c\r\n")
var newLine = []byte("\r\n")

var mode = &serial.Mode{
	BaudRate: 9600,
	DataBits: 8,
	StopBits: serial.OneStopBit,
	Parity:   serial.NoParity,
}

// DeskCycleSerial wrapper for a ReadWriteCloser port with an open desk cycle serial device.
// Provides methods for getting all desk cycle messages.
type DeskCycleSerial struct {
	port *respreader.ReadWriteCloser
	mu sync.Mutex
}

// Handshake returns the handshake response from the desk cycle
func (dc *DeskCycleSerial) Handshake() (string, error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return readLine(dc.port, handshakeMsg, 3)
}

// CurrentSpeed returns the current speed of the desk cycle
func (dc *DeskCycleSerial) CurrentSpeed() (float64, error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return readFloat(dc.port, speedMsg, 3)
}

// CurrentCadence returns the current cadence of the desk cycle
func (dc *DeskCycleSerial) CurrentCadence() (float64, error) {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	return readFloat(dc.port, cadenceMsg, 3)
}

// Close closes the desk cycle serial device
func (dc *DeskCycleSerial) Close() error {
	return dc.port.Close()
}

// NewDeskCycleSerial search serial ports for a desk cycle.
// Return DeskCycleSerial with the opened port
func NewDeskCycleSerial() (dc *DeskCycleSerial, err error) {
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

		response, err := readLine(portTimeout, handshakeMsg, 3)
		if err != nil {
			continue
		}

		if response == deskCycleName {
			dc = &DeskCycleSerial{port: portTimeout}
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

// readLine helper function writes a message and then reads from the port until a new line is found.
// It will retry sending/reading until maxAttempts is reached.
func readLine(port *respreader.ReadWriteCloser, message []byte, maxAttempts int) (string, error) {
	// initial message sent
	_, err := port.Write(message)
	if err != nil {
		return "", err
	}

	buff := make([]byte, 128)
	respBuilder := strings.Builder{}
	m := 0
	for attempt := 0; attempt < maxAttempts; {
		n, err := port.Read(buff)
		if err != nil {
			attempt += 1
			// if we haven't received any bytes yet retry sending the message
			if m == 0 {
				_, err = port.Write(message)
				if err != nil {
					return "", err
				}
			}
		}

		// if the response contains a newline return the response
		respBuilder.Write(buff[0:n])
		if bytes.Contains(buff, newLine) {
			return respBuilder.String(), nil
		}

		m += n
	}

	return "", errors.New("max attempts exceeded")
}

func readFloat(port *respreader.ReadWriteCloser, message []byte, maxAttempts int) (float64, error) {
	response, err := readLine(port, message, maxAttempts)

	if err != nil {
		return 0.0, err
	}

	return strconv.ParseFloat(strings.ReplaceAll(response, newLineStr, ""), 64)
}