package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func preview(fname string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", fname)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", fname)
	case "darwin":
		cmd = exec.Command("open", fname)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// use Start() so it doesn’t block while the viewer opens.
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start preview: %w", err)
	}

	// wait a short moment to help ensure the file opens properly
	time.Sleep(1 * time.Second)

	return nil
}
