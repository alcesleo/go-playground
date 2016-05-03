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
	fmt.Println(len(foundPrimes), "primes found")
	fmt.Println("finished in", end.Sub(start))
}

func primes(limit int) []int {
	var bitfields []bool = make([]bool, limit)
	var primes []int = []int{2} // 2 is the only even prime

	for num := 3; num < limit; num += 2 {
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
