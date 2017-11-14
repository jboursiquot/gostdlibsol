package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm root.")
}

type greetHandler struct{}

func (ph *greetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Gopher!")
}

func main() {
	gh := &greetHandler{}
	http.Handle("/greet/", gh)

	ph := newProverbsHandler()
	http.Handle("/proverbs/", ph)

	http.HandleFunc("/", handler)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
