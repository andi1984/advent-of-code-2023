package main

import (
	"bufio"
	Lib "day3/quizlib"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allDigitIndices [][]int
	var wholeMapping []string
	for scanner.Scan() {
		unparsedLine := scanner.Text()

		// Regexp number and non-digit groups (exclude dots)
		indices := Lib.GetDigitIndicesRegex(unparsedLine)
		allDigitIndices = append(allDigitIndices, indices)
		wholeMapping = append(wholeMapping, unparsedLine)
	}

	// Iterate over the digits

	hasAdjacent := func(x int, y int) bool {
		check := false
		if x > 0 {

			if !check && y > 0 {
				check = Lib.CharIsSymbol(wholeMapping[x-1][y-1])
			}

			if !check {
				check = Lib.CharIsSymbol(wholeMapping[x-1][y])
			}

			if !check && y < (len(wholeMapping[0])-1) {
				check = Lib.CharIsSymbol(wholeMapping[x-1][y+1])
			}
		}

		if !check && y > 0 {
			check = Lib.CharIsSymbol(wholeMapping[x][y-1])
		}

		if !check && y < (len(wholeMapping[0])-1) {
			check = Lib.CharIsSymbol(wholeMapping[x][y+1])
		}

		if !check && x < (len(wholeMapping)-1) {

			if !check && y > 0 {
				check = Lib.CharIsSymbol(wholeMapping[x+1][y-1])
			}

			if !check {
				check = Lib.CharIsSymbol(wholeMapping[x+1][y])
			}

			if !check && y < (len(wholeMapping[0])-1) {
				check = Lib.CharIsSymbol(wholeMapping[x+1][y+1])
			}
		}

		return check
	}
	isRuneADigit := func(r rune) bool { return unicode.IsDigit(r) }
	isRuneNotADigit := func(r rune) bool { return !isRuneADigit(r) }
	getNumber := func(x int, y int) int {
		input := wholeMapping[x]

		// find the index of the dot before the number
		startIndex := strings.LastIndexFunc(input[:y], isRuneNotADigit)
		var start int
		if startIndex == -1 {
			start = -1
		} else {
			start = startIndex - 1
		}

		// find the index of the dot after the number
		endIndex := strings.IndexFunc(input[y:], isRuneNotADigit)
		var end int
		if endIndex == -1 {
			end = y
		} else {
			end = y + endIndex
		}

		// get the substring that contains the number
		numStr := input[start+1 : end]
		// fmt.Println(x, y, numStr)

		filteredString := Lib.Filter(numStr, unicode.IsDigit)
		// convert the substring to an integer
		num, err := strconv.Atoi(filteredString)
		if err != nil {
			// handle error
			panic(err)
		}
		// print the result

		return num
	}

	sum := 0
	for x := 0; x < len(allDigitIndices)-1; x++ {
		lastOneIsAdjacent := false
		for i := 0; i < len(allDigitIndices[x])-1; i++ {
			y := allDigitIndices[x][i]
			if lastOneIsAdjacent {
				// Check if the next entry in allDigitIndices is +1 of previous
				// fmt.Println(y, allDigitIndices[x][i-1])
				if y == allDigitIndices[x][i-1]+1 {
					// we skip
					continue
				}
			}

			// fmt.Println(x, y)
			lastOneIsAdjacent = hasAdjacent(x, y)

			// Check if for a digit group there is a special char adjacent
			if lastOneIsAdjacent {
				//Get the full number at x,y
				// If yes, add it to the sum
				num := getNumber(x, y)
				// fmt.Println(num)
				sum = sum + num
			}
		}
	}

	fmt.Println(sum)
}
