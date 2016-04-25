package roman

import "strings"

type Pair struct {
	arabic int
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

func Convert(n int) string {
	roman := ""
	remainder := n
	for _, pair := range table {
		times := remainder / pair.arabic
		remainder = remainder % pair.arabic

		roman += strings.Repeat(pair.roman, times)
	}
	return roman
}
