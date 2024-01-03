package quizlib

import (
	"regexp"
	"unicode"
)

func GetDigitIndicesRegex(line string) []int {
	r := regexp.MustCompile(`\d`)
	matches := r.FindAllStringIndex(line, -1)

	var digitIndices []int
	for _, match := range matches {
		digitIndices = append(digitIndices, match[0])
	}

	return digitIndices
}

// filter returns a new string that contains only the characters of s that satisfy the predicate f
func Filter(s string, f func(rune) bool) string {
	var result []rune     // create an empty slice of runes
	for _, r := range s { // loop over the characters of s
		if f(r) { // if the character satisfies the predicate
			result = append(result, r) // append it to the result slice
		}
	}
	return string(result) // convert the slice back to a string and return it
}

func CharIsSymbol(char byte) bool {
	return !unicode.IsDigit(rune(char)) && char != '.'
}
