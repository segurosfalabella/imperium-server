package main

import (
	"net/http"

	"github.com/segurosfalabella/imperium-server/app"
)

type handler struct {
	port int
}

func (h *handler) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handler)
}

func (h *handler) ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}

func main() {
	h := new(handler)
	app.Start(h)
}
