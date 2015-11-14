package main

import (
	"fmt"
	"github.com/masonforest/slackbot"
)

type Stroopy struct {
	command string
}

func (s Stroopy) Respond(slashCommand slackbot.SlashCommand) string {
	return fmt.Sprintf("Hello, %s, I'm %s", slashCommand.SlashCommandData.Text, s.command)
}

func main() {
	slackbotServer := slackbot.Server{}
	slackbotServer.AddBot(Stroopy{command: "stroopy"})
	slackbotServer.Boot()
}
