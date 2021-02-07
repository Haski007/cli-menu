package main

import (
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

type Item string

func (i Item) String() string {
	return string(i)
}

const (
	FavVideosBot Item = "fav-videos-bot"


	Quit Item = "Quit"
)

func main()  {
	prompt := promptui.Select{
		Label:             "Choose bot session",
		Items:             []string{FavVideosBot.String(), Quit.String()},
		Size:              0,
		CursorPos:         0,
		IsVimMode:         false,
		HideHelp:          false,
		HideSelected:      false,
		Templates:         nil,
		Keys:              nil,
		Searcher:          nil,
		StartInSearchMode: false,
		Pointer:           nil,
		Stdin:             nil,
		Stdout:            nil,
	}

	_, result, err := prompt.Run()

	if err != nil {
		logrus.Fatalf("Prompt failed %s\n", err)
	}

	switch Item(result) {
	case Quit:
		os.Exit(0)
	default:
		cmd := exec.Command("zsh", "-c", "tmux attach -t " + result)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			logrus.Fatalf("[cmd.Run] err: %s", err)
		}
	}
}
