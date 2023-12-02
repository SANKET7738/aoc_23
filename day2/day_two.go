package day_two

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Problem struct{}

type Game struct {
	id        int
	max_red   int
	max_green int
	max_blue  int
}

type Set struct {
	red   int
	green int
	blue  int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (p Problem) ParseGameId(input string) int {
	temp := strings.Split(input, " ")
	id, _ := strconv.Atoi(temp[1])
	return id
}

func (p Problem) ParseSet(input string) Set {
	temp := strings.Split(input, ",")
	set := Set{}
	for _, colorStr := range temp {
		temp := strings.Split(colorStr, " ")
		count, _ := strconv.Atoi(temp[1])
		color := temp[2]
		if color == "red" {
			set.red = count
		} else if color == "green" {
			set.green = count
		} else if color == "blue" {
			set.blue = count
		}
	}
	return set
}

func (p Problem) ParseGameSets(input string, game *Game) {
	sets := strings.Split(input, ";")
	for _, set := range sets {
		parsedSet := p.ParseSet(set)
		game.max_red = max(game.max_red, parsedSet.red)
		game.max_green = max(game.max_green, parsedSet.green)
		game.max_blue = max(game.max_blue, parsedSet.blue)
	}
}

func (p Problem) isValidGame(game *Game) bool {
	if game.max_red > 12 || game.max_green > 13 || game.max_blue > 14 {
		return false
	}
	return true
}

func (p Problem) getProduct(game *Game) int {
	return game.max_red * game.max_green * game.max_blue
}

func (p Problem) Solve() {
	filepath := filepath.Join("day2", "input.txt")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	scanner := bufio.NewScanner(file)
	res := 0
	for scanner.Scan() {
		inputs := scanner.Text()
		temp := strings.Split(inputs, ":")
		game := Game{}
		game.id = p.ParseGameId(temp[0])
		p.ParseGameSets(temp[1], &game)
		val := p.getProduct(&game)
		res += val

	}
	fmt.Println(res)
}
