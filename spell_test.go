package spell_test

import (
	"github.com/daniellowtw/go-spell-number"
	"math"
	"testing"
)

func TestNegativeNumberReturnsEmptyString(t *testing.T) {
	actual := spell.Number(-1)
	if actual != "" {
		t.Errorf("Expected the empty string for negative numbers but got %s", actual)
		t.Fail()
	}
}

func TestLargestUInt64Number(t *testing.T) {
	expected := "eighteen quintillion four hundred forty six quadrillion seven hundred forty four trillion seventy three billion seven hundred nine million five hundred fifty one thousand six hundred fifteen"
	var x uint64 = math.MaxUint64
	actual := spell.LargeNumber(x)
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestLargestInt64Number(t *testing.T) {
	expected := "nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven"
	actual := spell.Number(math.MaxInt64)
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestLargeNumberSmokeTest(t *testing.T) {
	expected := "nine trillion one hundred twenty three billion four hundred fifty six million seven hundred eighty nine thousand one hundred twenty three"
	actual := spell.LargeNumber(9123456789123)
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestNumberProxySmokeTest(t *testing.T) {
	expected := "twelve"
	actual := spell.Number(12)
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
		t.Fail()
	}
}

func TestTrickyNumbers(t *testing.T) {
	tests := []struct {
		expected string
		number   uint64
	}{
		{
			expected: "zero",
			number:   0,
		},
		{
			expected: "ten",
			number:   10,
		},
		{
			expected: "twelve",
			number:   12,
		},
		{
			expected: "twenty",
			number:   20,
		},
		{
			expected: "twenty two",
			number:   22,
		},
		{
			expected: "one hundred",
			number:   100,
		},
		{
			expected: "one hundred eleven",
			number:   111,
		},
		{
			expected: "one thousand one hundred",
			number:   1100,
		},
		{
			expected: "one hundred eleven thousand one hundred eleven",
			number:   111111,
		},
	}
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			expected := test.expected
			actual := spell.LargeNumber(test.number)
			if actual != expected {
				t.Errorf("Expected %s but got %s", expected, actual)
				t.Fail()
			}

		})
	}
}

func TestShortFormNumberEnding(t *testing.T) {
	tests := []struct {
		expected string
		number   uint64
	}{
		{
			expected: "two thousand",
			number:   2 * 1e3,
		},
		{
			expected: "two million",
			number:   2 * 1e6,
		},
		{
			expected: "two billion",
			number:   2 * 1e9,
		},
		{
			expected: "two trillion",
			number:   2 * 1e12,
		},
		{
			expected: "two quadrillion",
			number:   2 * 1e15,
		},
		{
			expected: "two quintillion",
			number:   2 * 1e18,
		},
	}
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			expected := test.expected
			actual := spell.LargeNumber(test.number)
			if actual != expected {
				t.Errorf("Expected %s but got %s", expected, actual)
				t.Fail()
			}

		})
	}
}
