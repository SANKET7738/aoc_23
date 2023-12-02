package main

import (
	day_two "SANKET7738/aoc_23/day2"
)

type Problem interface {
	Solve()
}

func main() {
	var problem Problem
	problem = day_two.Problem{}
	problem.Solve()
}
