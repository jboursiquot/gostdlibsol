package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type key int

const usernameKey key = 22

func addUsername(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := r.Header.Get("X-Username")
		if u == "" {
			u = "Anonymous"
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, usernameKey, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func deepOperation(ctx context.Context) {
	u := ctx.Value(usernameKey).(string)
	log.Printf("[%s] Deep operation performed\n", u)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(usernameKey).(string)
	log.Printf("[%s] Handling greeting request", u)
	defer log.Printf("[%s] Handled greeting request", u)
	deepOperation(r.Context())
	w.Header().Set("X-Username", u)
	fmt.Fprintln(w, "Hello Gopher!")
}

func proverbHandler(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(usernameKey).(string)
	log.Printf("[%s] Handling proverb request", u)
	defer log.Printf("[%s] Handled proverb request", u)
	deepOperation(r.Context())
	w.Header().Set("X-Username", u)
	fmt.Fprintln(w, "Don't panic.")
}

func main() {
	http.HandleFunc("/greet", addUsername(greetHandler))
	http.HandleFunc("/proverb", addUsername(proverbHandler))
	log.Println("Staring server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
