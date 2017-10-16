package main

import (
	"fmt"
	"time"
)

var version string

func main() {
	fmt.Println("Hello Gophers!")
	fmt.Printf("Current time: %v\n", time.Now())
	//fmt.Println(person{"John", "Doe"})
	fmt.Printf("Version: %s\n", version)
}
