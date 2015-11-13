package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/masonforest/slackbot/Godeps/_workspace/src/github.com/goji/param"
	"github.com/masonforest/slackbot/Godeps/_workspace/src/github.com/zenazn/goji"
)

type SlashCommand struct {
	Token       string `param:"token"`
	TeamId      string `param:"team_id"`
	TeamDomain  string `param:"team_domain"`
	ChannelId   string `param:"channel_id"`
	ChannelName string `param:"channel_name"`
	UserId      string `param:"user_id"`
	UserName    string `param:"user_name"`
	Command     string `param:"command"`
	Text        string `param:"text"`
	ResponseUrl string `param:"response_url"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "", 400)
		return
	}

	var slashCommand SlashCommand

	err = param.Parse(r.PostForm, &slashCommand)
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	io.WriteString(w, fmt.Sprintf("%#v", slashCommand.Text))
}

func main() {
	flag.Set("bind", fmt.Sprint(":", os.Getenv("PORT")))
	goji.Post("/", hello)
	goji.Serve()
}
