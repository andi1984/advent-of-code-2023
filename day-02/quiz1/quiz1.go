package main

import (
	"bufio"
	Lib "day2/quizlib"
	"log"
	"os"
	"strconv"
	"strings"
)

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
			red := Lib.Get_Red(v)
			blue := Lib.Get_Blue(v)
			green := Lib.Get_Green(v)

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
