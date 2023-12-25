package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func get_red(foo string) int {
	re_red := regexp.MustCompile(`(\d\d?)\sred`)
	matches := re_red.FindSubmatch([]byte(foo))
	var red []byte
	for _, match := range matches {
		red = match
	}

	int, err := strconv.Atoi(string(red))
	if err != nil {
		return 0
	}

	return int
}

func get_blue(foo string) int {
	re_red := regexp.MustCompile(`(\d\d?)\sblue`)
	matches := re_red.FindSubmatch([]byte(foo))
	var red []byte
	for _, match := range matches {
		red = match
	}

	int, err := strconv.Atoi(string(red))
	if err != nil {
		return 0
	}

	return int
}

func get_green(foo string) int {
	re_red := regexp.MustCompile(`(\d\d?)\sgreen`)
	matches := re_red.FindSubmatch([]byte(foo))
	var red []byte
	for _, match := range matches {
		red = match
	}

	int, err := strconv.Atoi(string(red))
	if err != nil {
		return 0
	}

	return int
}

func main() {
	var RED_MAX = 12
	var GREEN_MAX = 13
	var BLUE_MAX = 14

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		unparsedLine := scanner.Text()
		gamesplit := strings.Split(unparsedLine, ":")
		runs_string := gamesplit[len(gamesplit)-1]
		game := gamesplit[0]

		game_number_as_string := strings.Split(game, " ")[1]
		gamenumber, _ := strconv.Atoi(game_number_as_string)
		runs := strings.Split(runs_string, ";")
		fulfills_all := true
		for _, v := range runs {
			if !fulfills_all {
				break
			}

			red := get_red(v)
			blue := get_blue(v)
			green := get_green(v)

			if red > RED_MAX || blue > BLUE_MAX || green > GREEN_MAX {
				fulfills_all = false
				break
			}

		}

		if fulfills_all {
			sum = sum + gamenumber
		}
	}

	println(sum)
}
