package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type book struct {
	title string
	path  string
	hist  histogram
}

type histogram struct {
	chars map[rune]int
}

var books = []*book{
	&book{title: "The Iliad", path: "../../data/the-iliad.txt"},
	&book{title: "The Underground Railroad", path: "../../data/the-underground-railroad.txt"},
	&book{title: "Pride and Prejudice", path: "../../data/pride-and-prejudice.txt"},
	&book{title: "The Republic", path: "../../data/the-republic.txt"},
	&book{title: "My Bondage and My Freedom", path: "../../data/my-bondage-and-my-freedom.txt"},
	&book{title: "War and Peace", path: "../../data/war-and-peace.txt"},
	&book{title: "Moby Dick", path: "../../data/moby-dick.txt"},
	&book{title: "Meditations", path: "../../data/meditations.txt"},
}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("Starting...")

	for _, b := range books {

		go func(b *book) { // https://golang.org/doc/faq#closures_and_goroutines
			log.Printf("Processing %s...", b.title)
			b.hist.chars = make(map[rune]int)

			file, err := os.Open(b.path)
			if err != nil {
				log.Fatalln(err)
			}

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				for _, c := range scanner.Text() {
					b.hist.chars[c]++
				}
			}

			file.Close()
			log.Printf("Done with %s", b.title)

		}(b)

	}

	tally(books...)
}

func tally(books ...*book) {
	log.Printf("Tallying results for %d books...", len(books))
	hist := histogram{chars: make(map[rune]int)}
	for _, book := range books {
		for key := range book.hist.chars {
			hist.chars[key] += book.hist.chars[key]
		}
	}
	printHist(&hist)
}

func printHist(h *histogram) {
	for k := range h.chars {
		fmt.Printf("%q=%d, ", k, h.chars[k])
	}
}
