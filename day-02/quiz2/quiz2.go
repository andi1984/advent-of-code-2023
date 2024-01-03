package main

import (
	"bufio"
	Lib "day2/quizlib"
	"log"
	"os"
	"strings"
)

func main() {
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

		runs := strings.Split(runs_string, ";")
		RED_MAX := 1
		BLUE_MAX := 1
		GREEN_MAX := 1
		for _, v := range runs {

			red := Lib.Get_Red(v)
			blue := Lib.Get_Blue(v)
			green := Lib.Get_Green(v)

			if RED_MAX < red {
				RED_MAX = red
			}

			if BLUE_MAX < blue {
				BLUE_MAX = blue
			}

			if GREEN_MAX < green {
				GREEN_MAX = green
			}

		}

		sum = sum + RED_MAX*BLUE_MAX*GREEN_MAX
	}

	println(sum)
}
