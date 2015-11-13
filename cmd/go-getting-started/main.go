package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func respond(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func main() {
	flag.Set("bind", fmt.Sprint(":", os.Getenv("PORT")))
	goji.Post("/", respond)
	goji.Serve()
}
