package day_four

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Problem struct{}

// {
// 	"1" : 1, = 1
// 	"2"=: 1, + 1 = 2
// 	"3": 1, + 1, + 2 = 4
// 	"4": 1, + 1, + 2, +4 = 8
// 	"5": 1, + 1 , + 4 , + 8, = 14
// 	"6": 1,  1
// }

func (p Problem) Solve() {
	filepath := filepath.Join("day4", "input.txt")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	scanner := bufio.NewScanner(file)
	cards := make([][]string, 0)
	cardFreqMap := make(map[int]int, 1)
	count := 0
	for scanner.Scan() {
		input := scanner.Text()
		temp := strings.Split(input, ":")[1]
		digits := strings.Split(temp, " | ")
		digitsSlice := strings.Split(digits[0], " ")
		digitsSlice = append(digitsSlice, strings.Split(digits[1], " ")...)
		cards = append(cards, digitsSlice)
		cardFreqMap[count] = 1
		count++
	}

	for idx, cardRow := range cards {
		count := 0
		digitsMap := make(map[int]int)
		for _, card := range cardRow {
			if card == "" {
				continue
			}
			digitInt, _ := strconv.Atoi(card)
			if _, ok := digitsMap[digitInt]; ok {
				digitsMap[digitInt]++
				count++
			} else {
				digitsMap[digitInt] = 1
			}
		}
		freq, _ := cardFreqMap[idx]
		for i := idx + 1; i <= idx+count; i++ {
			cardFreqMap[i] += freq
		}

	}

	res := 0
	for _, freq := range cardFreqMap {
		res += freq
	}
	fmt.Println(res)
}
