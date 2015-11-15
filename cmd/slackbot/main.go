package main

import (
	"github.com/masonforest/slackbot"
  "fmt"
  //"time"
)

func respond(r slackbot.Request) slackbot.Response {
  go func() {
    time.Sleep(1 * time.Second)
    r.Respond(slackbot.Response{Text: fmt.Sprintf("Hello %s",r.Data.Text)})
  }()
  return slackbot.EmptyResponse
}

func main() {
	server := slackbot.NewServer()
	server.AddCommand("/stroopy", respond)
	server.Boot()
}
