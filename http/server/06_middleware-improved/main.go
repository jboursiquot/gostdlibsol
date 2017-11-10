package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var requestsServed uint64

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
	log.Println("GREETED")
}

type statsHandler struct{}

func (sh *statsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greetings served: %d\n", atomic.LoadUint64(&requestsServed))
	log.Println("STATS PROVIDED")
}

func counter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		atomic.AddUint64(&requestsServed, 1)
		log.Println("COUNTER >> Counted")
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("LOGGER >> START %s %q\n", r.Method, r.URL.String())
		t := time.Now()
		log.Printf("LOGGER >> END %s %q (%v)\n", r.Method, r.URL.String(), time.Now().Sub(t))
		next.ServeHTTP(w, r)
	})
}

func use(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

func main() {

	// Before `use`, we wrap functions on functions. Cumbersome.
	// http.Handle("/greet", logger(counter(greetingHandler)))

	// With `use`, we provide a list of middleware to apply to handler. Preferred.
	http.Handle("/greet", use(http.HandlerFunc(greetHandler), counter, logger))

	sh := &statsHandler{}
	http.Handle("/stats", use(sh, logger))

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
