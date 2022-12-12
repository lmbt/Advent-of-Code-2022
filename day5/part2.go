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

	data, err := os.ReadFile("./input.txt")
	check(err)

	inputString := string(data)

	lines := strings.Split(inputString, "\n\n")
	setup := strings.Split(lines[0], "\n")
	// Remove extra text we don't need
	setup = setup[:len(setup)-1]
	numBays := ((len(setup[0]) - 3) / 4) + 1
	moves := strings.Split(lines[1], "\n")

	// Parse setup
	var bays [][]string

	for len(bays) < numBays {
		bays = append(bays, []string{})
	}

	for index := range setup {
		boxesData := setup[len(setup)-1-index]
		for i := 0; i < numBays; i++ {
			if string(boxesData[1+(i*4)]) != " " {
				bays[i] = append(bays[i], string(boxesData[1+(i*4)]))
			}
		}
	}

	// Iterate over movelist
	for _, move := range moves {
		instructions := strings.Split(move, " ")
		count, err := strconv.Atoi(instructions[1])
		check(err)
		source, err := strconv.Atoi(instructions[3])
		check(err)
		source--
		target, err := strconv.Atoi(instructions[5])
		check(err)
		target--
		//fmt.Printf("Move %d items from bay %d to bay %d\n", count, source, target)

		// Grab chunk of boxes
		sourceBayLen := len(bays[source])
		sourceBayChunkStart := sourceBayLen - count
		//fmt.Printf("Source bay length: %d | Chunk from %d to %d\n", sourceBayLen, sourceBayChunkStart, sourceBayChunkStart+count)
		sourceChunk := bays[source][sourceBayChunkStart:sourceBayLen]

		// Append chunk of boxes to target
		for _, box := range sourceChunk {
			bays[target] = append(bays[target], box)
		}

		// Remove chunk of boxes from source
		bays[source] = bays[source][:sourceBayChunkStart]

	}

	for _, bay := range bays {
		fmt.Print(bay[len(bay)-1])
	}
	fmt.Println()
}
