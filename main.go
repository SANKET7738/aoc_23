package main

import day_one "SANKET7738/aoc_23/day1"

type Problem interface {
	Solve()
}

func main() {
	var problem Problem
	problem = day_one.Problem{}
	problem.Solve()
}
