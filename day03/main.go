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

type Count struct {
	zeroes int64
	ones   int64
}

func mostPopular(lines []string) []Count {
	var counts []Count

	for r := 0; r < len(lines[0]); r++ {
		count := Count{zeroes: 0, ones: 0}

		for c := 0; c < len(lines); c++ {
			num, _ := strconv.Atoi(strings.Split(lines[c], "")[r])
			switch num {
			case 0:
				count.zeroes++
			case 1:
				count.ones++
			}
		}

		counts = append(counts, count)
	}

	return counts
}

func part1(lines []string) {
	rowWidth := len(lines[0])
	gamma := make([]string, rowWidth)
	epsilon := make([]string, rowWidth)

	counts := mostPopular(lines)
	for i, count := range counts {
		if count.ones > count.zeroes {
			gamma[i] = "1"
			epsilon[i] = "0"
		} else {
			gamma[i] = "0"
			epsilon[i] = "1"
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
