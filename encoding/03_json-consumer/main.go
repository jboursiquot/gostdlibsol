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
	filename := path.Join("..", "proverbs.json")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var proverbs []Proverb

	dec := json.NewDecoder(file)
	if err := dec.Decode(&proverbs); err != nil {
		log.Fatalln(err)
	}

	for _, p := range proverbs {
		log.Printf("%#v", p)
	}
}
