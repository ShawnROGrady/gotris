package main

import (
	"log"
	"os/exec"
	"runtime"
)

func setupTerm() (func() error, error) {
	var (
		limitCmd          *exec.Cmd
		removeEchoCmd     *exec.Cmd
		undoRemoveEchoCmd *exec.Cmd
	)

	switch runtime.GOOS {
	case "linux":
		limitCmd = exec.Command("stty", "-F", "/dev/tty", "-icanon", "min", "1")
		removeEchoCmd = exec.Command("stty", "-F", "/dev/tty", "-echo")
		undoRemoveEchoCmd = exec.Command("stty", "-F", "/dev/tty", "echo")
	case "darwin":
		limitCmd = exec.Command("stty", "--file", "/dev/tty", "-icanon", "min", "1")
		removeEchoCmd = exec.Command("stty", "--file", "/dev/tty", "-echo")
		undoRemoveEchoCmd = exec.Command("stty", "--file", "/dev/tty", "echo")
	default:
		limitCmd = exec.Command("stty", "--file", "/dev/tty", "-icanon", "min", "1")
		removeEchoCmd = exec.Command("stty", "--file", "/dev/tty", "-echo")
		undoRemoveEchoCmd = exec.Command("stty", "--file", "/dev/tty", "echo")
	}

	// set min number of characters for reading to 1
	if err := limitCmd.Run(); err != nil {
		log.Printf("error limiting input minimum: %v[%T]", err, err)
		return func() error { return nil }, err
	}
	// do not echo user input
	if err := removeEchoCmd.Run(); err != nil {
		log.Printf("error disabling user input echoing: %s", err)
		return func() error { return nil }, err
	}
	// re-enable echoing user input
	return undoRemoveEchoCmd.Run, nil
}
