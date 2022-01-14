package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

type Item string

func (i Item) String() string {
	return string(i)
}

const (
	crmBotTheSequel Item = "crm-bot-the-sequel"
	favVideosBot    Item = "fav-videos-bot"

	createAllSessions Item = "Create all sessions"
	killAllSessions   Item = "Kill all sessions"

	Quit Item = "Quit"
)

func main() {
	prompt := promptui.Select{
		Label: "Choose bot session",
		Items: []string{crmBotTheSequel.String(), favVideosBot.String(), killAllSessions.String(), createAllSessions.String(), Quit.String()},
	}

	var exit bool
	for !exit {
		_, result, err := prompt.Run()

		if err != nil {
			logrus.Fatalf("Prompt failed %s\n", err)
		}

		switch Item(result) {
		case createAllSessions:
			items, ok := prompt.Items.([]string)
			if !ok {
				logrus.Fatalf("prompt.Items is strings array")
			}

			// ---> parse only sessions names
			var sessions []string
			for _, item := range items {
				if strings.Contains(item, "-") {
					sessions = append(sessions, item)
				}
			}

			if err := CreateSessions(sessions); err != nil {
				logrus.Fatalf("create sessions err: %s", err)
			}
			exit = false

		case killAllSessions:
			if err := KillAllSessions(); err != nil {
				logrus.Errorf("kill all sessions err: %s", err)
			}

			logrus.Info("Killed all sessions")
			exit = false
		case Quit:
			exit = true
		default:
			cmd := exec.Command("zsh", "-c", "tmux attach -t "+result)
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				logrus.Fatalf("[cmd.Run] err: %s", err)
			}
			exit = true
		}
	}
}
