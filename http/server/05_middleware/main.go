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
	fmt.Fprintf(w, "Requests Served: %d\n", atomic.LoadUint64(&requestsServed))
	log.Println("STATS PROVIDED")
}

func counter(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		atomic.AddUint64(&requestsServed, 1)
		log.Println("COUNTER >> Counted")
	}
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("LOGGER >> START %s %q\n", r.Method, r.URL.String())
		t := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("LOGGER >> END %s %q (%v)\n", r.Method, r.URL.String(), time.Now().Sub(t))
	})
}

func main() {
	http.Handle("/greet", logger(counter(greetHandler)))

	sh := &statsHandler{}
	http.Handle("/stats", logger(sh))

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

// Two styles of middleware because of the following:
// http.Handle("/", http.HandlerFunc(f)) equivalent to
// http.HandleFunc("/", f)
// if f has signature func(http.ResponseWriter, *http.Response)

func middlewareUsingHandlerFunc(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// the middlerware's logic here...
		f(w, r) // equivalent to f.ServeHTTP(w, r)
	}
}

func middlewareUsingHander(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// the middlerware's logic here...
		next.ServeHTTP(w, r)
	})
}
