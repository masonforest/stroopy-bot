package slackbot

import (
	"github.com/masonforest/slackbot"
	"github.com/masonforest/slackbot/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
  "time"
)

import "testing"

func respond(r slackbot.Request) slackbot.Response {
	return slackbot.Response{Text: r.Data.Text}
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

	assert.Equal(t, string(response), "{\"text\":\"test_message\\n\"}")
}


func respondAsync(r slackbot.Request) slackbot.Response {
  go func() {
    r.Respond(slackbot.Response{Text: "pong"})
  }()

	return slackbot.EmptyResponse
}
func TestAsyncResponse(t *testing.T) {
  done := make(chan bool)
	server := slackbot.NewServer()
	server.AddCommand("/ping", respondAsync)

	responseServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    response, err := ioutil.ReadAll(r.Body)

    if err != nil {
      t.Fail()
    }

	  assert.Equal(t, string(response), "{\"text\":\"pong\"}")
    done <- true
  }))
	ts := httptest.NewServer(server)
	defer ts.Close()

	data := url.Values{"command": {"/ping"}, "response_url": {responseServer.URL}}

	res, err := http.PostForm(ts.URL, data)
	if err != nil {
		log.Fatal(err)
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	assert.Equal(t, string(response), "")


  select {
  case <-done:
  case <-time.After(time.Second * 1):
      t.Error(`Timeout`)
  }
}
