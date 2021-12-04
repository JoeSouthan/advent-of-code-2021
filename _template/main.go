package main

import (
	utils "advent-of-code/utils/files"
	"fmt"
	"os"
)

func part1(lines []string) {
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
	fmt.Println("Part 2")
	part2(lines)
}
