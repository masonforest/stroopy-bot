package main

import (
	"github.com/masonforest/slackbot"
  "fmt"
)

func respond(r slackbot.Request) slackbot.Response {
	return slackbot.Response{Text: fmt.Sprintf("Hello %s",r.Data.Text)}
}

func main() {
	server := slackbot.NewServer()
	server.AddCommand("/stroopy", respond)
	server.Boot()
}
