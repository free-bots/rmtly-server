package services

import (
	"fmt"
	"os/exec"
)

func SendAsync(title string, message string) {
	go sendMessage(title, message)
}

func sendMessage(title string, message string) {
	command := exec.Command("notify-send", fmt.Sprintf("%s", title), fmt.Sprintf("%s", message))
	err := command.Start()
	if err != nil {
		fmt.Println(err)
	}
}
