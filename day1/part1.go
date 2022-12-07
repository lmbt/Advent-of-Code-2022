package main

import (
	"fmt"
	"os"
	"sort"
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

	elfData := []int{}
	elves := 0

	for _, val := range lines {

		if len(elfData) <= elves {
			elfData = append(elfData, 0)
		}

		if val != "" {
			intVal, err := strconv.Atoi(val)
			check(err)

			elfData[elves] = elfData[elves] + intVal

		} else {

			elves++

		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfData)))

	fmt.Println(elfData[0])

}
