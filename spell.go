package spell

import (
	"strings"
)

var (
	oneToTen         = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	elevenToNineteen = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	multipleOfTen    = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	// Using the short scale as specified https://en.wikipedia.org/wiki/Names_of_large_numbers
	endingName = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion"}
)

// Number returns the spelling of the number. It returns the empty string if it is negative.
func Number(x int) string {
	if x < 0 {
		return ""
	}
	return LargeNumber(uint64(x))
}

// LargeNumber returns the spelling of the number.
func LargeNumber(x uint64) string {
	if x == 0 {
		return "zero"
	}
	var reverseStack []string
	suffixIdx := 0
	for x > 0 {
		spell3Digits := upTo1000(x % 1000)
		if spell3Digits == "" {
			suffixIdx++
			x /= 1000
			continue
		}

		suffix := endingName[suffixIdx]
		if suffix != "" {
			reverseStack = append(reverseStack, suffix)
		}
		reverseStack = append(reverseStack, spell3Digits)

		suffixIdx++
		x /= 1000
	}
	reverse(reverseStack)
	return strings.Join(reverseStack, " ")
}

func upTo1000(x uint64) string {
	remaining := upTo100(x % 100)
	if x < 100 {
		return remaining
	}
	hundreds := oneToTen[x/100]
	if x%100 == 0 {
		return hundreds + " hundred"
	}
	return strings.Join([]string{hundreds, "hundred", remaining}, " ")
}

func upTo100(x uint64) string {
	if x < 10 {
		return oneToTen[x]
	}
	if x < 20 {
		return elevenToNineteen[x-10]
	}

	tens := multipleOfTen[(x/10)%10]
	if x%10 == 0 {
		return tens
	}
	remainder := oneToTen[x%10]
	return tens + " " + remainder
}

func reverse(reverseStack []string) {
	for i := 0; i < len(reverseStack)/2; i++ {
		j := len(reverseStack) - 1 - i
		reverseStack[i], reverseStack[j] = reverseStack[j], reverseStack[i]
	}
}
