package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) ([]string, error) {
	var lines []string
	f, err := os.Open(path)
	if err != nil {
		return lines, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		lines = append(lines, s)
	}

	return lines, scanner.Err()
}

func main() {
	path := os.Args[1]
	lines, err := readFile(path)
	if err != nil {
		panic(err)
	}

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
