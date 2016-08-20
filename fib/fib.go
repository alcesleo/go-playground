package main

import "fmt"

func fib() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	nextFib, limit := fib(), 50
	for i := 0; i < limit; i++ {
		fmt.Println(nextFib())
	}
}
