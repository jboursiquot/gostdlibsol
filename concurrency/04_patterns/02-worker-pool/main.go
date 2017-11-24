package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

type book struct {
	title string
	path  string
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

type job struct {
	book *book
}

type result struct {
	job  job
	hist *histogram
}

func worker(g *sync.WaitGroup) {
	for job := range jobsChan {
		r := result{job, buildHistogram(job.book)}
		resultsChan <- r
	}
	g.Done()
}

func setupWorkerPool(numWorkers int) {
	var g sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		g.Add(1)
		go worker(&g)
	}
	g.Wait()
	close(resultsChan)
}

var jobsChan = make(chan job, 2)
var resultsChan = make(chan result, 2)

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("Starting...")

	go func() {
		for _, b := range books {
			jobsChan <- job{book: b}
		}
		close(jobsChan)
	}()

	done := make(chan bool)
	go func() {
		for r := range resultsChan {
			// printHist(r.job.book, r.hist)
			log.Printf("Job for %s done", r.job.book.title)
		}
		done <- true
	}()

	setupWorkerPool(2)
	<-done
	log.Println("Done")
}

func buildHistogram(b *book) *histogram {
	log.Printf("buildHistogram - Processing %s...", b.title)
	h := histogram{chars: make(map[rune]int)}

	file, err := os.Open(b.path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		for _, c := range scanner.Text() {
			h.chars[c]++
		}
	}

	return &h
}

func printHist(b *book, h *histogram) {
	log.Printf("printHist - %s", b.title)
	for k := range h.chars {
		fmt.Printf("%q=%d, ", k, h.chars[k])
	}
}
