package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// contains()
	// indexFunc()
	// checkStringEquality()
	// checkByteEquality()
	// checkCaseInsensitiveEquality()
	// compareBytes()
	// trimSpaces()
	// trimFunc()
	// replacer()
	// split()
	// fields()
	// join()
	// reader()
}

func contains() {
	fmt.Printf("%v\n", strings.Contains("hello world", "gophers"))
	fmt.Printf("%v\n", bytes.Contains([]byte("abc"), []byte("b")))
}

func indexFunc() {
	a := "This character ϵ is in this string"
	b := "This string does not contain epsilon"

	g := func(c rune) bool {
		return unicode.Is(unicode.Greek, c)
	}

	fmt.Println(strings.IndexFunc(a, g))
	fmt.Println(strings.IndexFunc(b, g))
}

func checkStringEquality() {
	fmt.Printf("%v\n", "a" == "b")
	fmt.Printf("%v\n", strings.ToUpper("a") == "A")
	fmt.Printf("%v\n", strings.ToUpper("ϵ") == "Ε")
}

func checkByteEquality() {
	fmt.Printf("%v\n", bytes.Equal([]byte{'a', 'b'}, []byte("ab")))
}

func checkCaseInsensitiveEquality() {
	fmt.Printf("%v\n", strings.EqualFold("ϵ", "Ε"))
}

func compareBytes() {
	fmt.Printf("%v\n", bytes.Compare([]byte{'a', 'b'}, []byte("ab")))
	fmt.Printf("%v\n", bytes.Compare([]byte{'a', 'b'}, []byte("abc")))
	fmt.Printf("%v\n", bytes.Compare([]byte{'a', 'b', 'c'}, []byte("ab")))
}

func trimSpaces() {
	fmt.Printf("%q\n", strings.TrimSpace(" hello "))
	fmt.Printf("%v\n", bytes.TrimSpace([]byte(" hello ")))
	fmt.Printf("%q\n", bytes.TrimSpace([]byte(" hello ")))
}

func trimFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	fmt.Printf("[%q]\n", strings.TrimFunc("   1234gophers unite5678", f))
}

func replacer() {
	r := strings.NewReplacer("alpha", "Α", "theta", "ϴ", "delta", "Δ")
	fmt.Printf("%q\n", r.Replace("The alpha differs from the theta which differs from the delta."))
}

func split() {
	alphabet := "alpha|beta|gamma"

	fmt.Printf("%q\n", strings.Split(alphabet, "|"))
	fmt.Printf("%q\n", bytes.Split([]byte(alphabet), []byte("|")))

	fmt.Printf("%q\n", strings.SplitAfter(alphabet, "|"))
	fmt.Printf("%q\n", bytes.SplitAfter([]byte(alphabet), []byte("|")))

	fmt.Printf("%q\n", strings.SplitN(alphabet, "|", 1))
	fmt.Printf("%q\n", bytes.SplitN([]byte(alphabet), []byte("|"), 2))
}

func fields() {
	alphabet := " alpha  beta     gamma "
	fmt.Printf("Fields are: %q\n", strings.Fields(alphabet))
	fmt.Printf("Fields are: %q\n", bytes.Fields([]byte(alphabet)))
}

func join() {
	alphabet := " alpha  beta     gamma "
	fields := strings.Fields(alphabet)
	fmt.Printf("Fields: %q\n", fields)
	fmt.Printf("Joined: %q\n", strings.Join(fields, "|"))
}

func reader() {
	var a, b, c string
	s := "a b c"
	r := strings.NewReader(s)
	_, err := fmt.Fscan(r, &a, &b, &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("a: %q\nb: %q\nc: %q\n", a, b, c)
}
