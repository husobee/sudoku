package sudoku_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/husobee/sudoku"
)

const (
	goodPuzzle string = `1 _ 3 _ _ 6 _ 8 _
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
_ 4 _ _ 7 8 _ 1 _
6 _ 8 _ _ 2 _ 4 _
_ 1 2 _ 4 5 _ 7 8
`
	NanPuzzle string = `a b c _ _ 6 _ 8 _
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
_ 4 _ _ 7 8 _ 1 _
6 _ 8 _ _ 2 _ 4 _
_ 1 2 _ 4 5 _ 7 8
`
	BadSpacesPuzzle string = `123   _ _ 6 _ 8 _
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
_ 4 _ _ 7 8 _ 1 _
6 _ 8 _ _ 2 _ 4 _
_ 1 2 _ 4 5 _ 7 8
`
	InvalidLengthPuzzle string = `1 2 3 _ _ 6 _ 8 _ 1 2 2
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
_ 4 _ _ 7 8 _ 1 _
6 _ 8 _ _ 2 _ 4 _
_ 1 2 _ 4 5 _ 7 8
`
	InvalidRowsToManyPuzzle string = `1 2 3 _ _ 6 _ 8 _
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
_ 4 _ _ 7 8 _ 1 _
6 _ 8 _ _ 2 _ 4 _
_ 1 2 _ 4 5 _ 7 8
_ 1 2 _ 4 5 _ 7 8
`
	InvalidRowsToFewPuzzle string = `1 2 3 _ _ 6 _ 8 _
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
`
)

func TestDump(t *testing.T) {
	p, err := sudoku.ParsePuzzle(strings.NewReader(goodPuzzle))
	if err != nil {
		t.Errorf("failed to parse a good puzzle, err=%s", err.Error())
	}
	buf := bytes.NewBuffer([]byte{})
	p.Dump(buf)
	if buf.String() != goodPuzzle {
		t.Errorf("failed to dump puzzle correctly")
	}
}
func TestParsePuzzle(t *testing.T) {
	if _, err := sudoku.ParsePuzzle(strings.NewReader(goodPuzzle)); err != nil {
		t.Errorf("failed to parse a good puzzle, err=%s", err.Error())
	}
	if _, err := sudoku.ParsePuzzle(strings.NewReader(NanPuzzle)); err == nil {
		t.Errorf("failed to error on a bad puzzle")
	}
	if _, err := sudoku.ParsePuzzle(strings.NewReader(BadSpacesPuzzle)); err == nil {
		t.Errorf("failed to error on a bad puzzle")
	}
	if _, err := sudoku.ParsePuzzle(strings.NewReader(InvalidRowsToFewPuzzle)); err == nil {
		t.Errorf("failed to error on a bad puzzle")
	}
	if _, err := sudoku.ParsePuzzle(strings.NewReader(InvalidRowsToManyPuzzle)); err == nil {
		t.Errorf("failed to error on a bad puzzle")
	}
	if _, err := sudoku.ParsePuzzle(strings.NewReader(InvalidLengthPuzzle)); err == nil {
		t.Errorf("failed to error on a bad puzzle")
	}
}

func BenchmarkParsePuzzle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sudoku.ParsePuzzle(strings.NewReader(goodPuzzle))
	}

}
