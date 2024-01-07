package util

import (
	"bytes"
	"errors"
	"strconv"
	"strings"

	"github.com/simpleiot/simpleiot/respreader"
)

var newLine = []byte("\r\n")

const newLineStr = "\r\n"

// readLine helper function writes a message and then reads from the port until a new line is found.
// It will retry sending/reading until maxAttempts is reached.
func ReadLine(port *respreader.ReadWriteCloser, message []byte, maxAttempts int) (string, error) {
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

func ReadFloat(port *respreader.ReadWriteCloser, message []byte, maxAttempts int) (float64, error) {
	response, err := ReadLine(port, message, maxAttempts)

	if err != nil {
		return 0.0, err
	}

	return strconv.ParseFloat(strings.ReplaceAll(response, newLineStr, ""), 64)
}
