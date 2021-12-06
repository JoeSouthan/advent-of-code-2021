package main

import (
	utils "advent-of-code/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Square struct {
	number int
	marked bool
}

type Board struct {
	lines [][]Square
}

func (b Board) printBoard() {
	for _, l := range b.lines {
		fmt.Println(l)
	}
}

func (s Square) mark() Square {
	s.marked = true
	return s
}

func markNumber(b Board, n int) Board {
	for i, l := range b.lines {
		for j, s := range l {
			if s.number == n {
				b.lines[i][j] = s.mark()

				fmt.Println("found", n, i, j)
			}
		}
	}

	return b
}

func (b Board) bingo() bool {
	return false
}

func parseLine(line string) []Square {
	l := strings.Split(line, " ")
	arr := make([]Square, len(l))

	for i, num := range l {
		n, _ := strconv.Atoi(num)
		arr[i] = Square{number: n, marked: false}
	}
	return arr
}

func countBoards(lines []string) int {
	count := 0
	for _, l := range lines {
		if l == "" {
			count++
		}
	}

	return count
}

func parseChoices(line string) []int {
	split := strings.Split(line, ",")
	var arr []int
	for _, s := range split {
		n, _ := strconv.Atoi(s)
		arr = append(arr, n)
	}
	return arr
}

func buildBoards(lines []string) []Board {
	boards := make([]Board, countBoards(lines))
	boardCount := 0

	for _, line := range lines {
		currentBoard := boards[boardCount]
		if line != "" {
			currentBoard.lines = append(currentBoard.lines, parseLine(line))
			boards[boardCount] = currentBoard
		} else {
			boardCount++
		}
	}

	return boards
}

func part1(lines []string) {
	choices := parseChoices(lines[0])
	lines = append(lines[2:], lines[:2]...)
	boards := buildBoards(lines)

	for _, c := range choices {
		for i, b := range boards {
			boards[i] = markNumber(b, c)
			if boards[i].bingo() {
				fmt.Println("bingo!")
				boards[i].printBoard()
				os.Exit(0)
			}
		}
	}

	fmt.Println(boards)
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
