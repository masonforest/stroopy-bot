package stroopybot

import (
	"github.com/masonforest/stroopybot"
	"github.com/masonforest/stroopybot/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestMain(t *testing.T) {
	stroopybot.SetupServer()

	ts := httptest.NewServer(stroopybot.Server)
	defer ts.Close()

	data := url.Values{"text": {"help"}, "command": {"/stroopy"}}

	res, err := http.PostForm(ts.URL, data)
	if err != nil {
		log.Fatal(err)
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	assert.Equal(t, string(response), "{\"response_type\":\"ephemeral\",\"text\":\"How can I help?\"}")
}
