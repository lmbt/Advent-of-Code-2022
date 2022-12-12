package main

import (
	"fmt"
	"os"
	"strconv"
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

	problemPairs := 0

	for _, line := range lines {

		sections := strings.Split(line, ",")

		sectionA := strings.Split(sections[0], "-")
		aStart, err := strconv.Atoi(sectionA[0])
		check(err)
		aEnd, err := strconv.Atoi(sectionA[1])
		check(err)

		sectionB := strings.Split(sections[1], "-")
		bStart, err := strconv.Atoi(sectionB[0])
		check(err)
		bEnd, err := strconv.Atoi(sectionB[1])
		check(err)

		// If A fully contains B
		if aStart <= bStart && aEnd >= bEnd {
			//fmt.Printf(" A contains B at line %d\n", index)
			problemPairs++

			// If B fully contains A
		} else if aStart >= bStart && aEnd <= bEnd {
			//fmt.Printf(" B contains A at line %d\n", index)
			problemPairs++
		}

	}

	fmt.Println(problemPairs)
}
