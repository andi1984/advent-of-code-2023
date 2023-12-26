package quizlib

import (
	"regexp"
	"strconv"
)

func Get_Red(foo string) int {
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

func Get_Blue(foo string) int {
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

func Get_Green(foo string) int {
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
