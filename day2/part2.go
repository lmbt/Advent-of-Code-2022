package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("./input.txt")
	check(err)

	inputString := string(dat)

	lines := strings.Split(inputString, "\n")

	score := 0

	for _, line := range lines {
		plays := strings.Split(line, " ")

		// Determine what to play
		if plays[1] == "X" {
			// Need to lose so only record score for your play
			if plays[0] == "A" {
				score += 3
			} else if plays[0] == "B" {
				score += 1
			} else {
				score += 2
			}
		} else if plays[1] == "Y" {
			// Need to draw so record score for play plus 3
			if plays[0] == "A" {
				score += 4
			} else if plays[0] == "B" {
				score += 5
			} else {
				score += 6
			}
		} else {
			// Need to win so record score for play plus 6
			if plays[0] == "A" {
				score += 8
			} else if plays[0] == "B" {
				score += 9
			} else {
				score += 7
			}
		}

	}

	fmt.Println(score)

}
