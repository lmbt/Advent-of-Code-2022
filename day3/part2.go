package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Source: problem description
// a-z is 1-26
// A-Z is 27-52
func findPriorityVal(letter rune) int {
	var value int

	if unicode.IsUpper(letter) {
		value = int(letter - 38)
	} else {
		value = int(letter - 96)
	}

	return value
}

// True if rune slice contains specified rune
func runeSliceContains(needle rune, haystack []rune) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}
	return false
}

func main() {

	dat, err := os.ReadFile("./input.txt")
	check(err)

	inputString := string(dat)

	lines := strings.Split(inputString, "\n")

	prioritySum := 0

	for lineIndex, line := range lines {
		// Prevent duplicate character counting
		var usedRunes []rune

		// Find badge ID for this elf group
		if (lineIndex+1)%3 == 0 {

			for _, letter := range line {

				// This rune must exist in this line and the previous 2 lines to be a match
				// Way too many characters on this line...
				if strings.Contains(lines[lineIndex-2], string(letter)) && strings.Contains(lines[lineIndex-1], string(letter)) && !runeSliceContains(letter, usedRunes) {
					usedRunes = append(usedRunes, letter)
					prioritySum += findPriorityVal(letter)
				}

			}

		}

	}

	fmt.Println(prioritySum)

}
