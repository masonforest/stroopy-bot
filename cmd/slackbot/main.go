package main

import (
	"github.com/masonforest/slackbot"
  "fmt"
)

func respond(r slackbot.Request) string {
  return fmt.Sprintf("Hello %s", r.Data.Text)
}

func main() {
	server := slackbot.NewServer()
	server.AddCommand("/stroopy", respond)
	server.Boot()
}
