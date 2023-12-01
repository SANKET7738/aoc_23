package day_one

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"unicode"
)

type Problem struct{}

var firstCharDigitMap = map[string][]string{
	"o": {"one"},
	"t": {"two", "three"},
	"f": {"four", "five"},
	"s": {"six", "seven"},
	"e": {"eight"},
	"n": {"nine"},
}

var lastCharDigitMap = map[string][]string{
	"e": {"one", "three", "five", "nine"},
	"o": {"zero", "two"},
	"r": {"four"},
	"x": {"six"},
	"n": {"seven"},
	"t": {"eight"},
}

type Result struct {
	doExist bool
	digit   string
}

type RequiredDigit struct {
	digit string
	isInt bool
}

func (p Problem) checkDigitExists(ch chan Result, wg *sync.WaitGroup, input string, idx int, digit string, reverse bool) {
	defer wg.Done()
	initialIdx := idx
	count := 0
	if reverse {
		for digitIdx := len(digit) - 1; digitIdx >= 0; digitIdx-- {
			if initialIdx < 0 || input[initialIdx] != digit[digitIdx] {
				break
			}
			count++
			initialIdx--
		}
	} else {
		for digitIdx := range digit {
			if initialIdx >= len(input) || input[initialIdx] != digit[digitIdx] {
				break
			}
			count++
			initialIdx++
		}
	}
	if count == len(digit) {
		result := Result{
			doExist: true,
			digit:   digit,
		}
		ch <- result
	}

	result := Result{
		doExist: false,
		digit:   digit,
	}
	ch <- result
}

func (p Problem) getContaintedDigit(input string, idx int, startingChar string, reverse bool) string {
	var digits []string
	if reverse {
		digits, _ = lastCharDigitMap[startingChar]
	} else {
		digits, _ = firstCharDigitMap[startingChar]
	}
	ch := make(chan Result)
	wg := &sync.WaitGroup{}

	for _, digit := range digits {
		wg.Add(1)
		go p.checkDigitExists(ch, wg, input, idx, digit, reverse)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		if result.doExist {
			return result.digit
		}
	}

	return ""
}

func (p Problem) getFirstDigit(input string) RequiredDigit {
	for idx, char := range input {
		if unicode.IsDigit(char) {
			return RequiredDigit{
				digit: string(char),
				isInt: true,
			}
		}
		_, exists := firstCharDigitMap[string(char)]
		if exists {
			digit := p.getContaintedDigit(input, idx, string(char), false)
			if digit != "" {
				return RequiredDigit{
					digit: digit,
					isInt: false,
				}
			}
		}
	}
	return RequiredDigit{}
}

func (p Problem) getLastDigit(input string) RequiredDigit {
	for idx := len(input) - 1; idx >= 0; idx-- {
		char := input[idx]
		if unicode.IsDigit(rune(char)) {
			return RequiredDigit{
				digit: string(char),
				isInt: true,
			}
		}
		_, exists := lastCharDigitMap[string(char)]
		if exists {
			digit := p.getContaintedDigit(input, idx, string(char), true)
			if digit != "" {
				return RequiredDigit{
					digit: digit,
					isInt: false,
				}
			}
		}
	}
	return RequiredDigit{}
}

func (p Problem) getFinalResult(firstDigit RequiredDigit, lastDigit RequiredDigit) int {
	if lastDigit.digit == "" {
		lastDigit = firstDigit
	}

	digitMap := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var firstDigitStr, lastDigitStr string
	if !firstDigit.isInt {
		firstDigitStr = string(digitMap[firstDigit.digit])
	} else {
		firstDigitStr = firstDigit.digit
	}

	if !lastDigit.isInt {
		lastDigitStr = string(digitMap[lastDigit.digit])
	} else {
		lastDigitStr = lastDigit.digit
	}

	res, _ := strconv.Atoi(firstDigitStr + lastDigitStr)
	return res
}

func (p Problem) Solve() {
	filepath := filepath.Join("day1", "input.txt")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := 0

	for scanner.Scan() {
		inputs := scanner.Text()
		var firstDigit, lastDigit RequiredDigit
		firstDigit = p.getFirstDigit(inputs)
		lastDigit = p.getLastDigit(inputs)
		res += p.getFinalResult(firstDigit, lastDigit)
	}
	fmt.Println(res)
}
