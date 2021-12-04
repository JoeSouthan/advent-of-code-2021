package main

import (
	utils "advent-of-code/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bitsToInt(s []string) int64 {
	r, _ := strconv.ParseInt(strings.Join(s[:], ""), 2, 64)
	return r
}

func part1(lines []string) {
	rowWidth := len(lines[0])
	noRows := len(lines)
	gamma := make([]string, rowWidth)
	epsilon := make([]string, rowWidth)

	for r := 0; r < rowWidth; r++ {
		ones, zeroes := 0, 0

		for c := 0; c < noRows; c++ {
			num, _ := strconv.Atoi(strings.Split(lines[c], "")[r])
			switch num {
			case 0:
				zeroes++
			case 1:
				ones++
			}
		}

		if ones > zeroes {
			gamma[r] = "1"
			epsilon[r] = "0"
		} else {
			gamma[r] = "0"
			epsilon[r] = "1"
		}
	}

	fmt.Println("epsilon", bitsToInt(epsilon))
	fmt.Println("gamma", bitsToInt(gamma))
	fmt.Println("consumption", bitsToInt(gamma)*bitsToInt(epsilon))
}

func part2(lines []string) {
}

func main() {
	path := os.Args[1]
	lines, err := utils.ParseInput(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1")
	part1(lines)
	// fmt.Println("Part 2")
	// part2(lines)
}
