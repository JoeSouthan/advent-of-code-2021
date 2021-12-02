package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(path string) ([]int, error) {
	var lines []int
	f, err := os.Open(path)
	if err != nil {
		return lines, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return lines, err
		}

		lines = append(lines, i)
	}

	return lines, scanner.Err()
}

func isIncrease(now int, previous int) bool {
	return now > previous
}

func main() {
	path := os.Args[1]
	lines, err := readFile(path)
	if err != nil {
		panic(err)
	}

	increases := 0

	for i := 0; i < len(lines)-3; i += 1 {
		if lines[i+3] > lines[i] {
			increases += 1
		}
	}
	fmt.Println(increases)
}
