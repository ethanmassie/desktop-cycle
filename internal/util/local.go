package util

import (
	"errors"
	"os"
	"runtime"
)

const LINUX_PATH = "/.local/share/desktop-cycle/"

// Get file path to the local/App Data directory for the GOOS
func GetLocalPath() (string, error) {
	// TODO: add cross platform paths
	home, _ := os.UserHomeDir()
	if runtime.GOOS == "linux" {
		return home + LINUX_PATH, nil
	}

	return "", errors.New("unsupported OS " + runtime.GOOS)
}
