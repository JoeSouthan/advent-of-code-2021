package main

import (
	utils "advent-of-code/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bitsToInt(s string) int64 {
	r, _ := strconv.ParseInt(s, 2, 64)
	return r
}

type Count struct {
	zeroes int64
	ones   int64
}

func (c Count) mostCommon() string {
	common := "1"
	if c.zeroes >= c.ones {
		common = "0"
	}

	return common
}

func (c Count) leastCommon() string {
	common := "0"
	if c.mostCommon() == "0" {
		common = "1"
	}

	return common
}

func countPositions(lines []string) []Count {
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

func filterValues(lines []string, f string, position int) []string {
	var filtered []string
	for _, line := range lines {
		l := strings.Split(line, "")
		if l[position] == f {
			filtered = append(filtered, line)
		}
	}

	return filtered
}

func oxygenRating(lines []string) string {
	for i := 0; i < len(lines[0]); i++ {
		counts := countPositions(lines)
		count := counts[i]
		if count.mostCommon() == "1" || count.zeroes == count.ones {
			lines = filterValues(lines, "1", i)
		} else {
			lines = filterValues(lines, "0", i)
		}
	}

	return lines[0]
}

func co2Rating(lines []string) string {
	for i := 0; i < len(lines[0]); i++ {
		counts := countPositions(lines)
		count := counts[i]
		if len(lines) == 1 {
			break
		}
		if count.leastCommon() == "0" || count.zeroes == count.ones {
			lines = filterValues(lines, "0", i)
		} else {
			lines = filterValues(lines, "1", i)
		}
	}

	return lines[0]
}

func part1(lines []string) {
	rowWidth := len(lines[0])
	gamma := make([]string, rowWidth)
	epsilon := make([]string, rowWidth)

	counts := countPositions(lines)
	for i, count := range counts {
		if count.ones > count.zeroes {
			gamma[i] = "1"
			epsilon[i] = "0"
		} else {
			gamma[i] = "0"
			epsilon[i] = "1"
		}
	}

	e := strings.Join(epsilon, "")
	g := strings.Join(gamma, "")

	fmt.Println("epsilon", bitsToInt(e))
	fmt.Println("gamma", bitsToInt(g))
	fmt.Println("consumption", bitsToInt(g)*bitsToInt(e))
}

func part2(lines []string) {
	o2 := bitsToInt(oxygenRating(lines))
	co2 := bitsToInt(co2Rating(lines))
	fmt.Println("o2", o2)
	fmt.Println("co2", co2)
	fmt.Println("consumption", o2*co2)
}

func main() {
	path := os.Args[1]
	lines, err := utils.ParseInput(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 2")
	part2(lines)
}
