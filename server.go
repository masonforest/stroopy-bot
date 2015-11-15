package stroopybot

import (
	"github.com/masonforest/stroopybot/Godeps/_workspace/src/github.com/masonforest/slackbot"
)

var Server *slackbot.Server

func SetupServer() {
	Server = slackbot.NewServer()
	Server.AddCommand("/stroopy", Respond)
}
func Boot() {
	SetupServer()
	Server.Boot()
}
