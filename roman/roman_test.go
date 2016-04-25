package roman

import (
	"testing"
)

func TestConversions(t *testing.T) {
	expectations := []struct {
		n        int    // input
		expected string // expected result
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},

		{501, "DI"},
		{1000, "M"},
		{2020, "MMXX"},
	}

	for _, test := range expectations {
		result := Convert(test.n)
		if result != test.expected {
			t.Errorf("Expected %v to be converted to %v, got %#v", test.n, test.expected, result)
		}
	}
}
