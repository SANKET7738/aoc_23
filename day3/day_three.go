package day_three

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

type Problem struct{}

// 0 0 1 2 3 4 5 6 7 8 9
// 0 4 6 7 . . 1 1 4 . .
// 1 . . . * . . . . . .
// 2 . . 3 5 . . 6 3 3 .
// 3 . . . . . . # . . .
// 4 6 1 7 * . . . . . .
// 5 . . . . . + . 5 8 .
// 6 . . 5 9 2 . . . . .
// 7 . . . . . . 7 5 5 .
// 8 . . . $ . * . . . .
// 9 . 6 6 4 . 5 9 8 . .

type Board [][]string

func (p Problem) isSpecialChar(r rune) bool {
	return !unicode.IsNumber(r) && !unicode.IsLetter(r)
}

func (p Problem) hasValidNeighbor(board Board, stars Board, row int, col int, digit string) bool {
	if row < 0 || row >= len(board) || col < 0 || col >= len((board)[0]) {
		return false
	}
	neighborChar := (board)[row][col]
	if neighborChar == "*" {
		stars[row][col] += digit
		stars[row][col] += "*"
	}
	if neighborChar != "." && p.isSpecialChar([]rune(neighborChar)[0]) {
		return true
	}
	return false
}

func (p Problem) IsDigitValid(board Board, stars Board, row int, col int, digit string) bool {
	var digitStart, digitEnd int
	if col-len(digit)-1 >= 0 {
		digitStart = col - len(digit) + 1
	} else {
		digitStart = 0
	}
	digitEnd = col

	// check one row up
	for i := digitStart - 1; i <= digitEnd+1; i++ {
		if p.hasValidNeighbor(board, stars, row-1, i, digit) {
			return true
		}
	}

	// check same row
	if p.hasValidNeighbor(board, stars, row, digitStart-1, digit) || p.hasValidNeighbor(board, stars, row, digitEnd+1, digit) {
		return true
	}

	// check one row down
	for i := digitStart - 1; i <= digitEnd+1; i++ {
		if p.hasValidNeighbor(board, stars, row+1, i, digit) {
			return true
		}
	}
	return false
}

func (p Problem) Solve() {
	filepath := filepath.Join("day3", "input.txt")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	scanner := bufio.NewScanner(file)
	board := Board{}
	for scanner.Scan() {
		row := make([]string, 0)
		for _, char := range scanner.Text() {
			row = append(row, string(char))
		}
		board = append(board, row)

	}

	stars := make(Board, len(board))
	for i := range stars {
		stars[i] = make([]string, len(board[i]))
		for j := range stars[i] {
			stars[i][j] = ""
		}
	}

	res := 0

	for i, row := range board {
		currentDigit := ""
		for j, char := range row {
			if unicode.IsDigit([]rune(char)[0]) {
				currentDigit += string(char)
			}
			if !unicode.IsDigit([]rune(char)[0]) && currentDigit != "" {
				var currentDigitInt int
				currentDigitInt, _ = strconv.Atoi(currentDigit)
				if p.IsDigitValid(board, stars, i, j-1, currentDigit) {
					res += currentDigitInt
				}
				currentDigit = ""
			}
		}
		if currentDigit != "" {
			if p.IsDigitValid(board, stars, i, len(row)-1, currentDigit) {
				var currentDigitInt int
				currentDigitInt, _ = strconv.Atoi(currentDigit)
				res += currentDigitInt
			}
		}
	}

	res = 0
	for _, row := range stars {
		for _, char := range row {
			if len(char) > 1 {
				temp := strings.Split(char, "*")
				if len(temp) > 2 {
					a, _ := strconv.Atoi(temp[0])
					b, _ := strconv.Atoi(temp[1])
					res += a * b
				}
			}

		}
	}
	fmt.Println(res)
}
