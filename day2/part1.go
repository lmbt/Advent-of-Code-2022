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

		// Add score for playing your move
		switch plays[1] {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		default:
			fmt.Println("Unknown instruction")
		}

		// Add score for outcome of round
		if plays[1] == "X" && plays[0] == "C" ||
			plays[1] == "Y" && plays[0] == "A" ||
			plays[1] == "Z" && plays[0] == "B" {
			// Winner
			score += 6
		} else if plays[1] == "X" && plays[0] == "A" ||
			// Draw
			plays[1] == "Y" && plays[0] == "B" ||
			plays[1] == "Z" && plays[0] == "C" {
			score += 3
		} else {
			// Loser
			score += 0
		}

	}

	fmt.Println(score)

}
