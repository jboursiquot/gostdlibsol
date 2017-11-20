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
	Text     string
	reviewed bool
}

func main() {
	proverbs := []Proverb{
		Proverb{ID: 1, Text: "Don't panic."},
		Proverb{ID: 2, Text: "Concurrency is not parallelism."},
		Proverb{ID: 3, Text: "Documentation is for users."},
		Proverb{ID: 4, Text: "The bigger the interface, the weaker the abstraction."},
		Proverb{ID: 5, Text: "Make the zero value useful."},
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
