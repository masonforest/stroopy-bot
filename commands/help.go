package commands

import (
	"github.com/masonforest/slackbot"
)

func Help(r slackbot.Request) slackbot.Response {
  return slackbot.Response{
    ResponseType: slackbot.EMPHEMERAL,
    Text: "How can I help?",
  }
}
