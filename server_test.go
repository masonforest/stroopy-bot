package slackbot

import (
	"github.com/masonforest/slackbot"
	"github.com/masonforest/slackbot/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
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
