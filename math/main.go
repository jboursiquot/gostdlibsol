package main

import (
	"fmt"
	"math/big"
	"math/cmplx"
	"math/rand"
	"time"
)

func main() {
	bigInt()
	// primeCheck()
	// bigFloat()
	// bigRat()
	// complexNums()
	// random()
}

func bigInt() {
	// n := 34500000000000000000 // Nope!
	n := new(big.Int)
	n.SetString("34500000000000000000", 10)
	fmt.Printf("n = %v - %T\n", n, n)
	fmt.Printf("%s\n", n.Text(10))

	m := big.NewInt(128)
	fmt.Printf("m = %v - %T\n", m, m)

	n.Add(n, m)
	fmt.Printf("n = %v - %T\n", n, n)

	o := new(big.Int).Mul(n, m)
	fmt.Printf("o = %v - %T\n", o, o)

	fmt.Printf("n.Cmp(o): %d\n", n.Cmp(o))
	fmt.Printf("n.Cmp(m): %d\n", n.Cmp(m))
	fmt.Printf("n.Cmp(n): %d\n", n.Cmp(n))
}

func primeCheck() {
	primes := []*big.Int{
		big.NewInt(329),
		big.NewInt(337),
		big.NewInt(347),
		big.NewInt(349),
		big.NewInt(350),
		big.NewInt(358),
	}
	for _, p := range primes {
		fmt.Printf("%v a prime? %t\n", p, p.ProbablyPrime(1))
	}
}

func bigFloat() {
	var pi big.Float
	pi.Parse("3.14159265358979323846264338327950288419716939937510582097494459", 10)
	fmt.Printf("pi = %.10g\n", &pi)
}

func bigRat() {
	n := big.NewRat(1, 2)
	m := big.NewRat(100, 200)
	fmt.Printf("n = %v/%v\n", n.Num(), n.Denom())
	fmt.Printf("m = %v/%v\n", m.Num(), m.Denom())
	fmt.Printf("n.Cmp(m) = %v\n\n", n.Cmp(m))

	o := new(big.Rat)
	_, _ = fmt.Sscan(".8", o)
	fmt.Printf("o = %v/%v\n", o.Num(), o.Denom())
	fmt.Printf("n.Cmp(o) = %v\n", n.Cmp(o))
}

func complexNums() {
	n := complex(1, 2) // using builtin complex(...)
	m := 3 + 6i        // using shorthand syntax

	sum := n + m
	abs := cmplx.Abs(sum)
	fmt.Printf("sum: %v, %T, Abs: %v, %T\n", sum, sum, abs, abs)

	prod := n * m
	icos := cmplx.Acos(prod)
	fmt.Printf("prod: %v, %T, Acos: %v, %T\n", prod, prod, icos, icos)

	pow := cmplx.Pow(n, m)
	fmt.Printf("pow: %v, %T\n", pow, pow)

	sqrt := cmplx.Sqrt(pow)
	fmt.Printf("sqrt: %v, %T\n", sqrt, sqrt)
}

func random() {
	proverbs := []string{
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Documentation is for users.",
		"Don't panic.",
	}

	n := rand.Intn(len(proverbs))
	fmt.Println("Random Proverb:", proverbs[n], n)

	rand.Seed(103)
	n = rand.Intn(len(proverbs))
	fmt.Println("Random Proverb:", proverbs[n], n)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n = r.Intn(len(proverbs))
	fmt.Println("Random Proverb:", proverbs[n], n)
}
