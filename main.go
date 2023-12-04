package main

import (
	day_three "SANKET7738/aoc_23/day3"
)

type Problem interface {
	Solve()
}

func main() {
	var problem Problem
	problem = day_three.Problem{}
	problem.Solve()
}
