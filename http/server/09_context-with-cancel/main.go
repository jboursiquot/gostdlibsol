package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling greeting request")
	defer log.Println("Handled greeting request")

	completeAfter := time.After(5 * time.Second)

	ctx, cancel := context.WithCancel(r.Context())
	time.AfterFunc(3*time.Second, func() {
		cancel()
	})

	for {
		select {
		case <-completeAfter:
			fmt.Fprintln(w, "Hello Gopher!")
			return
		case <-ctx.Done():
			err := ctx.Err()
			log.Printf("Context Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusRequestTimeout)
			return
		default:
			time.Sleep(1 * time.Second)
			log.Println("Greetings are hard. Thinking...")
		}
	}
}

func main() {
	http.HandleFunc("/", greetHandler)
	log.Println("Staring server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
