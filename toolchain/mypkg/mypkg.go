package mypkg

import "fmt"

// SayHello says hello
func SayHello(name string) {
	fmt.Printf("Hello %s!", name)
}

// hidden is a hidden struct.
type hidden struct{}
