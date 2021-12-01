package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(path string) ([]int64, error) {
	var lines []int64
	f, err := os.Open(path)

	if err != nil {
		return lines, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)

		if err != nil {
			return lines, err
		}

		lines = append(lines, i)
	}

	return lines, scanner.Err()
}

func isIncrease(now int64, previous int64) bool {
	return now > previous
}

func main() {
	lines, err := readFile("day01/input.txt")

	if err != nil {
		panic(err)
	}

	previous := lines[0]
	increases := 0

	for i := 0; i < len(lines); i++ {
		if isIncrease(lines[i], previous) {
			increases += 1
		}

		previous = lines[i]
	}

	fmt.Println(increases)
}
