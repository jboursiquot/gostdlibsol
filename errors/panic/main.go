package main

import "log"

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("boom() failed: %v\n", r)
		}
	}()

	log.Println("Hello Gopher!")
	boom()
}

func boom() {
	panic("BOOM!")
}
