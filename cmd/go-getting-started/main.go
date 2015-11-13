package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zenazn/goji"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong!\n")
}

func main() {
	flag.Set("bind", fmt.Sprint(":", os.Getenv("PORT")))
	goji.Post("/", hello)
	goji.Serve()
}
