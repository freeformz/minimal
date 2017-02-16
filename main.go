package main

import (
	"io"
	"net/http"
	"os"

	"github.com/heroku/slog"
)

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := slog.Context{}
	ctx["hello"] = "world"
	io.WriteString(w, ctx.String())
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5000"
	}
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
