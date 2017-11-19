package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

// Proverb states a general truth or piece of advice.
type Proverb struct {
	ID       int
	Value    string
	reviewed bool
}

func main() {
	proverbs := []Proverb{
		Proverb{ID: 1, Value: "Don't panic."},
		Proverb{ID: 2, Value: "Concurrency is not parallelism."},
		Proverb{ID: 3, Value: "Documentation is for users."},
		Proverb{ID: 4, Value: "The bigger the interface, the weaker the abstraction."},
		Proverb{ID: 5, Value: "Make the zero value useful."},
	}

	filename := path.Join("..", "proverbs.json")
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	if err := enc.Encode(proverbs); err != nil {
		log.Fatalln(err)
	}

	log.Println("Done.")
}
