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

func bookStage(books []*book) <-chan *book {
	out := make(chan *book)
	go func() {
		for _, b := range books {
			log.Printf("bookStage - %s", b.title)
			out <- b
		}
		close(out)
	}()
	return out
}

func histStage(in <-chan *book) <-chan *book {
	out := make(chan *book)
	go func() {
		for b := range in {
			log.Printf("histStage - Processing %s...", b.title)
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
			log.Printf("histStage - Done with %s", b.title)

			out <- b
		}
		close(out)
	}()
	return out
}

func tallyStage(h *histogram, in <-chan *book) <-chan *book {
	out := make(chan *book)
	go func() {
		for b := range in {
			log.Printf("tallyStage - %s", b.title)
			for key := range b.hist.chars {
				h.chars[key] += b.hist.chars[key]
			}
			out <- b
		}
		close(out)
	}()
	return out
}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("Starting...")

	hist := histogram{chars: make(map[rune]int)}

	bs := bookStage(books)
	hs := histStage(bs)
	ts := tallyStage(&hist, hs)

	for b := range ts {
		log.Printf("main - %s", b.title)
	}

	log.Println("main - HISTOGRAM")
	printHist(&hist)
}

func printHist(h *histogram) {
	for k := range h.chars {
		fmt.Printf("%q=%d, ", k, h.chars[k])
	}
}
