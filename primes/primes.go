package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	limit, err := strconv.Atoi(os.Args[1])
	if err != nil || limit < 0 {
		fmt.Println("Usage: primes <upto>")
		os.Exit(1)
	}

	start := time.Now()
	foundPrimes := primes(limit)
	end := time.Now()

	// Print the primes unless given a second argument
	if len(os.Args) == 2 {
		fmt.Println(foundPrimes)
	}
	fmt.Println("finished in", end.Sub(start))
}

func primes(limit int) []int {
	var bitfields []bool = make([]bool, limit)
	var primes []int

	for num := 2; num < limit; num++ {
		// Collect prime
		if !bitfields[num] {
			primes = append(primes, num)
		}

		// Cross off all multiples
		for multiple := num * 2; multiple < limit; multiple += num {
			bitfields[multiple] = true
		}
	}

	return primes
}
