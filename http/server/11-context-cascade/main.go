package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jboursiquot/things"
)

type opStatus int

const (
	opCompleted opStatus = iota
	opCanceled
	opTimedOut
)

func doSlowOperation(ctx context.Context, statusChan chan opStatus) {
	log.Println("doSlowOperation >> called...")

	d := 4 * time.Second
	completionTimeout := time.After(d)

	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if err == context.Canceled {
				log.Println("doSlowOperation >> context was canceled")
				statusChan <- opCanceled
				return
			} else if err == context.DeadlineExceeded {
				log.Println("doSlowOperation >> context deadline was exceeded")
				statusChan <- opTimedOut
				return
			}
		case <-completionTimeout:
			log.Printf("doSlowOperation >> %v passed, operation complete", d)
			statusChan <- opCompleted
			return
		default:
			log.Println("doSlowOperation >> slow operation...")
			time.Sleep(1 * time.Second)
		}
	}
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("greetHandler >> handling greeting request")
	defer log.Println("greetHandler >> handled greeting request")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	statusChan := make(chan opStatus)
	go doSlowOperation(ctx, statusChan)

	for {
		select {
		case status := <-statusChan:
			log.Println("greetHandler >> read from statusChan...")
			if status == opCanceled {
				log.Println("greetHandler >> operation was cancelled...")
				http.Error(w, "Internal request was cancelled", http.StatusInternalServerError)
				return
			}
			if status == opTimedOut {
				log.Println("greetHandler >> operation timedout...")
				http.Error(w, "Internal operation took too long", http.StatusRequestTimeout)
				return
			}
			log.Println("greetHandler >> responding to client...")
			fmt.Fprintln(w, "Hello Gopher!")
			return
		case <-things.DoHardThings(ctx):
			log.Printf("greetHandler >> things did hard thingss")
		}
	}
}

func main() {
	http.HandleFunc("/", greetHandler)
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
