
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

# Async Example
Sometimes commands take time to process. This example shows how you can repond
to commands asynchronously

    package main

    import (
      "github.com/masonforest/slackbot"
      "fmt"
      "time"
    )

    func respond(r slackbot.Request) slackbot.Response {
      go func() {
        time.Sleep(1 * time.Second)
        r.Respond(slackbot.Response{Text: fmt.Sprintf("Pong %s",r.Data.Text)})
      }()
      return slackbot.EmptyResponse
    }

    func main() {
      server := slackbot.NewServer()
      server.AddCommand("/slowping", respond)
      server.Boot()
    }
