package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: rect <number>")
		os.Exit(1)
	}
	num, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("rect only works with positive integers")
		os.Exit(1)
	}

	fmt.Println(rect(num))
}

func rect(num uint64) (uint64, uint64) {
	x := uint64(math.Sqrt(float64(num)))
	y := x

	for x <= num && y >= 1 {
		mult := x * y
		if mult == num {
			break;
		}

		if mult < num {
			x++
		} else {
			y--
		}
	}

	return x, y
}
