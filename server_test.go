package slackbot

import (
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
  "github.com/masonforest/slackbot"
)

import "testing"

func respond(r slackbot.Request) string {
  return r.Data.Text
}

func TestServer(t *testing.T) {
  server := slackbot.NewServer()
  server.AddCommand("/ping", respond)

  ts := httptest.NewServer(server)
  defer ts.Close()

  data := url.Values{"text": {"test_message\n"}, "command": {"/ping"}}

  res, err := http.PostForm(ts.URL, data)
  if err != nil {
    log.Fatal(err)
  }
  response, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  res.Body.Close()

  assert.Equal(t, string(response), "test_message\n")
}

