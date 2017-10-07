package main

import (
	"fmt"
	"strconv"
)

func main() {
	// stringToInt()
	// intToString()
	// stringToAltBaseInt()
	// formatInt()
	// appendInt()
	// parseFloat()
	// formatFloat()
	// parseBool()
	// quoting()
}

func stringToInt() {
	n, err := strconv.Atoi("345")
	// n, err := strconv.Atoi("34500000000000000000")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v - %T\n", n, n)
}

func intToString() {
	n := 345
	// n := 34500000000000000000
	s := strconv.Itoa(n)
	fmt.Printf("%v - %T\n", s, s)
}

func stringToAltBaseInt() {
	s := "1604"
	i, _ := strconv.ParseInt(s, 8, 0)
	fmt.Printf("%v - %T\n", i, i)
	// fmt.Printf("%o\n", i)
}

func formatInt() {
	x := int64(345)
	s := strconv.FormatInt(x, 10)
	fmt.Printf("%T, %v\n", s, s)

	b := int64(-1024)
	t := strconv.FormatInt(b, 2)
	fmt.Printf("%T, %v\n", t, t)
}

func appendInt() {
	b := []byte("What are we appending?: ")
	b = strconv.AppendInt(b, -345, 10)
	fmt.Printf("%v\n", string(b))
}

func parseFloat() {
	π := "3.1415926535"
	f, err := strconv.ParseFloat(π, 64)
	fmt.Printf("%v - %T - %v\n", f, f, err)
}

func formatFloat() {
	π := 3.1415926535
	f := strconv.FormatFloat(π, 'e', 2, 64)
	fmt.Printf("%v - %T\n", f, f)
}

func parseBool() {
	list := []string{"1", "t", "T", "TRUE", "true", "True", "Yes", "Yay", "Yippie", "0", "f", "F", "FALSE", "false", "False", "Nope", "Nah", "No"}
	for _, b := range list {
		b, err := strconv.ParseBool(b)
		fmt.Printf("%v - %T - %v\n", b, b, err)
	}
}

func quoting() {
	fmt.Println(strconv.Quote("There is a π symbol here as well as a tab	."))
	fmt.Println(strconv.QuoteToASCII("There is a π symbol here as well as a tab	."))
}
