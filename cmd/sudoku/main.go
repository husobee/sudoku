package main

import (
	"log"
	"os"

	"github.com/husobee/sudoku"
)

func main() {
	// take stdin and make a Puzzle
	p, err := sudoku.ParsePuzzle(os.Stdin)
	if err != nil {
		// bad input
		log.Fatalf("Invalid Puzzle: %s", err.Error())
	}
	// attempt to solve the puzzle
	if err := p.BacktrackSolve(); err != nil {
		// couldn't solve the puzzle
		log.Fatalf("Error solving Puzzle: %s", err.Error())
	}
	// dump out the solution to stdout
	p.Dump(os.Stdout)
}
