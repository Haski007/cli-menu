package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func CreateSessions(sessionsNames []string) error {
	for i, sessionName := range sessionsNames {
		cmd := exec.Command("zsh", "-c", "tmux  new -d -s " + sessionName)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("create session %s err: %w", sessionName, err)
		}
		logrus.Infof("created session #%d %s", i+1, sessionName)
	}

	return nil
}