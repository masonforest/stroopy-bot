package main

import (
	"fmt"
	"github.com/masonforest/stroopybot/Godeps/_workspace/src/github.com/masonforest/slackbot"
)

func respond(r slackbot.Request) slackbot.Response {
	return slackbot.Response{
		ResponseType: slackbot.IN_CHANNEL,
		Text:         fmt.Sprintf("Hello %s", r.Data.Text),
	}
}

func main() {
	server := slackbot.NewServer()
	server.AddCommand("/stroopy", respond)
	server.Boot()
}
