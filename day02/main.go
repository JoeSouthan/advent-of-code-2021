package main

import (
	utils "advent-of-code/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) {
	x := 0
	y := 0

	for i := 0; i < len(lines); i++ {
		op := strings.Split(lines[i], " ")
		direction := op[0]
		increment, err := strconv.Atoi(op[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			x += increment
		case "up":
			y += increment
		case "down":
			y -= increment
		default:
			fmt.Println("unknown")
		}
	}

	fmt.Println(x, y)
	fmt.Println(x * y)
}

func part2(lines []string) {
	x := 0
	y := 0
	aim := 0

	for i := 0; i < len(lines); i++ {
		op := strings.Split(lines[i], " ")
		direction := op[0]
		increment, err := strconv.Atoi(op[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			x += increment
			y += increment * aim
		case "up":
			aim -= increment
		case "down":
			aim += increment
		default:
			fmt.Println("unknown")
		}
	}

	fmt.Println(x, y, aim)
	fmt.Println(x * y)
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
