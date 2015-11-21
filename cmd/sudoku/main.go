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
		log.Fatalf("Invalid Puzzle: %s", err.Error())
	}
	p.Dump(os.Stdout)
}
