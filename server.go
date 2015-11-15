package slackbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/masonforest/slackbot/Godeps/_workspace/src/github.com/gorilla/schema"
	"io"
	"net/http"
	"os"
)

type RequestData struct {
	Token       string
	TeamId      string
	TeamDomain  string
	ChannelId   string
	ChannelName string
	UserId      string
	UserName    string
	Command     string
	Text        string
	ResponseUrl string
}

type Request struct {
	w    http.ResponseWriter
	r    *http.Request
	Data *RequestData
}

type Response struct {
	Text string
}

func (r Response) toString() string {
	data := map[string]string{"text": r.Text}
	s, _ := json.Marshal(data)
	return string(s)
}

func (s Request) Respond(response Response) {
	var byteString = []byte(response.toString())
	req, err := http.NewRequest("POST", s.Data.Text, bytes.NewBuffer(byteString))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

type Command interface {
	Respond(slashCommand Request) string
}

type Server struct {
	commands map[string]func(Request) string
}

func NewServer() *Server {
	return &Server{commands: make(map[string]func(Request) string)}
}
func (s *Server) AddCommand(name string, command func(Request) string) {
	s.commands[name] = command
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "", 400)
		return
	}

	data := &RequestData{}
	decoder := schema.NewDecoder()

	err = decoder.Decode(data, r.PostForm)
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	slashCommand := Request{w: w, r: r, Data: data}

	var response string
	c := s.commands[slashCommand.Data.Command]
	response = c(slashCommand)
	io.WriteString(w, response)
}

func (s Server) Boot() {
	http.HandleFunc("/", s.ServeHTTP)
	http.ListenAndServe(fmt.Sprint(":", os.Getenv("PORT")), nil)
}
