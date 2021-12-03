package utils

import (
	"bufio"
	"os"
)

func ParseInput(path string) ([]string, error) {
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
