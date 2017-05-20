package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

const envTerminalWidth = "COLUMNS"

func getTerminalWidth() (int, error) {
	if width, err := getUserConfiguredTerminalWidth(); err == nil {
		return width, nil
	}
	if width, _, err := terminal.GetSize(syscall.Stdout); err == nil {
		return width, nil
	}
	return 0, fmt.Errorf("could not determine terminal width")
}

func getUserConfiguredTerminalWidth() (int, error) {
	return strconv.Atoi(os.Getenv(envTerminalWidth))
}
