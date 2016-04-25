package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	arabic uint64
	roman  string
}

var table = []Pair{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func Convert(n uint64) string {
	roman := ""
	remainder := n
	for _, pair := range table {
		times := remainder / pair.arabic
		remainder = remainder % pair.arabic

		roman += strings.Repeat(pair.roman, int(times))
	}
	return roman
}

func main() {
	errMsg := "Usage: roman <integer>"
	if len(os.Args) != 2 {
		fmt.Println(errMsg)
		os.Exit(1)
	}

	input, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(errMsg)
		os.Exit(1)
	}

	fmt.Println(Convert(input))
}
