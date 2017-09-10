package main

import "fmt"

var (
	i = 100
	f = 3.14
	b = true
	s = "Clear is better than clever."
	p = struct{ x, y int64 }{252, 101}
)

var bi int64 = -922337203685
var ui uint64 = 1844679551615

func main() {
	outputToStdout()
}

func outputToStdout() {
	fmt.Println(i, f, b, s, bi, ui, p)
	fmt.Printf("i = %v\nf = %v\nb = %v\ns = %v\nbi = %v\nui = %v\np = %v\n", i, f, b, s, bi, ui, p)
	fmt.Printf("i = %#v\nf = %#v\nb = %#v\ns = %#v\nbi = %#v\nui = %#v\np = %#v\n", i, f, b, s, bi, ui, p)
	fmt.Printf("i = %T\nf = %T\nb = %T\ns = %T\nbi = %T\nui = %T\np = %T\n", i, f, b, s, bi, ui, p)
}
