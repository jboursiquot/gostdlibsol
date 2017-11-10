package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var requestsHandled uint64

func greetHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&requestsHandled, 1)

	log.Printf("START %s %q\n", r.Method, r.URL.String())
	t := time.Now()

	fmt.Fprintln(w, "Hello")
	log.Println("GREETED")

	log.Printf("END %s %q (%v)\n", r.Method, r.URL.String(), time.Now().Sub(t))
}

type statsHandler struct{}

func (sh *statsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("START %s %q\n", r.Method, r.URL.String())
	t := time.Now()

	fmt.Fprintf(w, "Requests Handled: %d\n", atomic.LoadUint64(&requestsHandled))
	log.Println("STATS PROVIDED")

	log.Printf("END %s %q (%v)\n", r.Method, r.URL.String(), time.Now().Sub(t))
}

func main() {
	http.HandleFunc("/greet", greetHandler)

	sh := &statsHandler{}
	http.Handle("/stats", sh)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
