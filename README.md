
# Slackbot library for golang

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Example

      package main

      import (
        "fmt"
        "github.com/masonforest/slackbot"
      )

      type SimpleBot struct {
        command string
      }

      func (s SimpleBot) Respond(slashCommand slackbot.SlashCommand) string {
        return fmt.Sprintf("Hello, %s, I'm %s", slashCommand.SlashCommandData.Text, s.command)
      }

      func main() {
        slackbotServer := slackbot.Server{}
        slackbotServer.AddBot(SimpleBot{command: "simplebot"})
        slackbotServer.Boot()
      }
