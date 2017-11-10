package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Server >> Request received...[%s] %s", r.Method, r.RequestURI)
		msg := "Hello Gopher!"
		log.Printf("Server >> Sending %q", msg)
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, msg)
	}))
	defer ts.Close()

	log.Printf("Client >> Making request to test server: %s", ts.URL)
	t := time.Now()
	r, err := http.Get(ts.URL)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	msg := strings.TrimSpace(string(b))
	log.Printf("Client >> Received response %q in %v", msg, time.Since(t))
}
