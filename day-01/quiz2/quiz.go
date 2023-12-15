package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		unparsedLine := scanner.Text()
		subLine := ""

		// First replace all word-based numbers into numbers from left to right iteratively
		for _, c := range unparsedLine {
			char := string(c)
			subLine = subLine + char

			// We have to make sure to keep overlapping characters for edging
			// numbers
			subLine = strings.NewReplacer(
				"zero", "z0o",
				"one", "o1e",
				"two", "t2o",
				"three", "t3e",
				"four", "f4r",
				"five", "f5e",
				"six", "s6x",
				"seven", "s7n",
				"eight", "e8t",
				"nine", "n9e",
			).Replace(subLine)
		}

		line := subLine

		// Iterate over chars
		var first *int
		var last *int
		for _, c := range line {
			char := string(c)
			if char >= "0" && char <= "9" {
				intValue, err := strconv.Atoi(char)
				if err != nil {
					break
				}

				if first == nil {
					first = &intValue
				} else {
					last = &intValue
				}
			}
		}

		if err != nil {
			log.Fatal(err)
		}

		if last == nil {
			last = *(&first)
		}

		value := 10*(*first) + *last
		// fmt.Printf("%v--->%v\n", line, value)

		sum = sum + value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum is: %v", sum)
}
