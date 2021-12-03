package main

import (
	utils "advent-of-code/utils/files"
	"fmt"
	"os"
)

func part1(lines []string) {
	increases := 0

	for i := 0; i < len(lines)-1; i += 1 {
		if lines[i+1] > lines[i] {
			increases += 1
		}
	}
	fmt.Println(increases)
}

func part2(lines []string) {
	increases := 0

	for i := 0; i < len(lines)-3; i += 1 {
		if lines[i+3] > lines[i] {
			increases += 1
		}
	}
	fmt.Println(increases)
}

func main() {
	path := os.Args[1]
	lines, err := utils.ParseInput(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1")
	part1(lines)
	fmt.Println("Part 1")
	part2(lines)
}
