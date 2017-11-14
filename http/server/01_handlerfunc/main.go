package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm root.")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
