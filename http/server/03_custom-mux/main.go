package main

import (
	"fmt"
	"log"
	"net/http"
)

// muxer is a custom request multiplexer
type muxer struct{}

func (m *muxer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/greet":
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello Gopher!")
		}(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	m := &muxer{}
	log.Println("Staring server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", m))
}
