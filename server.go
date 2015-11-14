package slackbot

import (
	"flag"
	"fmt"
	"github.com/masonforest/slackbot/Godeps/_workspace/src/github.com/goji/param"
	"github.com/masonforest/slackbot/Godeps/_workspace/src/github.com/zenazn/goji"
	"io"
	"net/http"
	"os"
)

type SlashCommandData struct {
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

type SlashCommand struct {
	HTTPResponeWriter http.ResponseWriter
	HTTPRequest       *http.Request
	SlashCommandData  SlashCommandData
}

type Bot interface {
	Respond(slashCommand SlashCommand) string
}

type Server struct {
	bot Bot
}

func (s *Server) AddBot(bot Bot) {
	s.bot = bot
}

func (s Server) Boot() {
	flag.Set("bind", fmt.Sprint(":", os.Getenv("PORT")))
	goji.Post("/", s.ResponseHandler)
	goji.Serve()
}

func (s Server) ResponseHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "", 400)
		return
	}

	var slashCommandData SlashCommandData

	slashCommand := SlashCommand{HTTPResponeWriter: w, HTTPRequest: r, SlashCommandData: slashCommandData}
	err = param.Parse(r.PostForm, &slashCommand.SlashCommandData)
	if err != nil {
		http.Error(slashCommand.HTTPResponeWriter, "", 500)
		return
	}

	response := s.bot.Respond(slashCommand)
	io.WriteString(w, response)
}
