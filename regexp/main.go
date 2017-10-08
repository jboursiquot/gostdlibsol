package main

import (
	"fmt"
	"regexp"
)

func main() {
	match()
	// find()
	// submatch()
	// replace()
	// split()
}

func match() {
	p := "(foo|faux)bar"

	s1 := "fallbar"
	match, err := regexp.MatchString(p, s1)
	fmt.Printf("MatchString - %v, %v\n", match, err)

	s2 := "foobar"
	match, err = regexp.MatchString(p, s2)
	fmt.Printf("MatchString - %v, %v\n", match, err)

	// p = "go(pher)*"

	// b1 := []byte("Cgo is not Go.")
	// match, err = regexp.Match(p, b1)
	// fmt.Printf("Match - %v, %v\n", match, err)

	// b2 := []byte("Where are all my gophers?")
	// match, err = regexp.Match(p, b2)
	// fmt.Printf("Match - %v, %v\n", match, err)
}

func find() {
	s1 := "Cgo is not Go."
	s2 := "Where are all my gophers?"
	r, _ := regexp.Compile("(g|G)o(pher)*")
	fmt.Printf("Pattern: %v\n", r.String())

	fmt.Printf("MatchString(%q): %v\n", s1, r.MatchString(s1))
	fmt.Printf("MatchString(%q): %v\n", s2, r.MatchString(s2))

	// fmt.Printf("FindString(%q): %v\n", s1, r.FindString(s1))
	// fmt.Printf("FindString(%q): %v\n", s2, r.FindString(s2))

	// fmt.Printf("FindAllString(%q): %v\n", s1, r.FindAllString(s1, -1))
	// fmt.Printf("FindAllString(%q): %v\n", s2, r.FindAllString(s2, -1))

	// fmt.Printf("FindStringIndex(%q): %v\n", s1, r.FindStringIndex(s1))
	// fmt.Printf("FindStringIndex(%q): %v\n", s2, r.FindStringIndex(s2))
}

func submatch() {
	s1 := "Cgo is not Go."
	s2 := "Where are all my gophers?"
	r, _ := regexp.Compile("([a-z]+)er")
	fmt.Printf("Pattern: %v\n", r.String())

	fmt.Printf("FindAllStringSubmatch(%q): %v\n", s1, r.FindAllStringSubmatch(s1, -1))
	fmt.Printf("FindAllStringSubmatch(%q): %v\n", s2, r.FindAllStringSubmatch(s2, -1))

	// fmt.Printf("FindSubmatch(%q): %s\n", s1, r.FindSubmatch([]byte(s1)))
	// fmt.Printf("FindSubmatch(%q): %s\n", s2, r.FindSubmatch([]byte(s2)))
}

func replace() {
	s1 := "Cgo is not Go."
	s2 := "Where are all my gophers?"
	s3 := "Clear is better than clever."
	r, _ := regexp.Compile("(g|G)o(pher)*")
	fmt.Printf("Pattern: %v\n", r.String())

	fmt.Printf("ReplaceAll - %v: %s\n", s1, r.ReplaceAll([]byte(s1), []byte("lol")))
	fmt.Printf("ReplaceAll - %v: %s\n", s2, r.ReplaceAll([]byte(s2), []byte("lol")))
	fmt.Printf("ReplaceAll - %v: %s\n", s3, r.ReplaceAll([]byte(s3), []byte("lol")))

	// fmt.Printf("ReplaceAllStringFunc - %v: %v\n", s1, r.ReplaceAllStringFunc(s1, strings.ToUpper))
	// fmt.Printf("ReplaceAllStringFunc - %v: %v\n", s2, r.ReplaceAllStringFunc(s2, strings.ToUpper))
	// fmt.Printf("ReplaceAllStringFunc - %v: %v\n", s3, r.ReplaceAllStringFunc(s3, strings.ToUpper))
}

func split() {
	s1 := "Cgo is not Go."
	s2 := "Where are all my gophers?"
	s3 := "Clear is better than clever."
	r, _ := regexp.Compile("(g|G)o(pher)*")
	fmt.Printf("Pattern: %v\n", r.String())

	fmt.Printf("Split - %v: %q\n", s1, r.Split(s1, -1))
	fmt.Printf("Split - %v: %q\n", s2, r.Split(s2, -1))
	fmt.Printf("Split - %v: %q\n", s3, r.Split(s3, -1))
}
