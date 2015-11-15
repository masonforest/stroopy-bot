package commands

import (
  "log"
  "fmt"
	"github.com/masonforest/stroopybot/Godeps/_workspace/src/github.com/masonforest/slackbot"
  "github.com/stellar/go-stellar-base/keypair"
)

func NewAddress(r slackbot.Request) slackbot.Response {
  kp, err := keypair.Random()

  if err != nil {
    log.Fatal(err)
  }

  r.Respond(slackbot.Response{
    ResponseType: slackbot.IN_CHANNEL,
    Text: fmt.Sprintf("%s generated a new stellar address: %s", r.Data.UserName, kp.Address()),
  })

  r.Respond(slackbot.Response{
    ResponseType: slackbot.EMPHEMERAL,
    Text: fmt.Sprintf("Here is your private key. Keep it safe. Anyone who has access to you private key can access your funds: %s", kp.(*keypair.Full).Seed()),
  })

  return slackbot.EmptyResponse
}
