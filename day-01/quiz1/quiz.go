package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		line := scanner.Text()
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

		sum = sum + value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum is: %v", sum)
}
