package stroopybot

import (
	"fmt"
	"github.com/masonforest/stroopybot/Godeps/_workspace/src/github.com/masonforest/slackbot"
	"github.com/masonforest/stroopybot/commands"
	"strings"
)

var cm = map[string]func(slackbot.Request) slackbot.Response{
	"help": commands.Help,
	"new":  commands.NewAddress,
}

func Respond(r slackbot.Request) slackbot.Response {
	s := strings.Split(r.Data.Text, " ")

	if command, present := cm[s[0]]; present {
		return command(r)
	} else {
		return slackbot.Response{
			ResponseType: slackbot.EMPHEMERAL,
			Text:         fmt.Sprintf("Sorry %s, I don't know how to %s", r.Data.UserName, s[0]),
		}
	}
}
