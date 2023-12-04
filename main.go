package main

import (
	day_four "SANKET7738/aoc_23/day4"
)

type Problem interface {
	Solve()
}

func main() {
	var problem Problem
	problem = day_four.Problem{}
	problem.Solve()
}
