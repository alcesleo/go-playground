package main

import (
	"fmt"
	"time"
)

func main() {
	limit := 1000

	start := time.Now()
	primes := primes(limit)
	end := time.Now()

	fmt.Println(primes)
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
