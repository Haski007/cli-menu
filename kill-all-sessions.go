package main

import (
	"fmt"
	"os"
	"os/exec"
)

func KillAllSessions() error {
	cmd := exec.Command("zsh", "-c", "tmux kill-server")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run cmd err: %w", err)
	}
	return nil
}
