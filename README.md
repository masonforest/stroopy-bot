
# Slackbot library for golang

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Example

      package main

      import (
        "github.com/masonforest/slackbot"
        "fmt"
      )

      func respond(r slackbot.Request) string {
        return fmt.Sprintf("Pong %s", r.Data.Text)
      }

      func main() {
        server := slackbot.NewServer()
        server.AddCommand("/ping", respond)
        server.Boot()
      }
