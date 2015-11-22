package sudoku

import (
	"bufio"
	"errors"
	"io"
)

const (
	Underscore        byte = 0x5f
	Space                  = 0x20
	Zero                   = 0x30
	One                    = 0x31
	Nine                   = 0x39
	Newline                = 0x0a
	MaxInputRowLength int  = 17
)

var (
	ErrInvalidNumber     = errors.New("invalid number")
	ErrInvalidCharacter  = errors.New("invalid character")
	ErrInvalidLineLength = errors.New("invalid line length")
	ErrInvalidRowCount   = errors.New("invalid number of rows")
	ErrStuck             = errors.New("stuck in backtrack")
	ErrNoSolution        = errors.New("no solution")
)

func isSpace(c byte) bool {
	return c == Space
}
func isBlank(c byte) bool {
	return c == Underscore
}

func isNumber(c byte) bool {
	if _, err := asciiToNumber(c); err != nil {
		return false
	}
	return true

}

func asciiToNumber(c byte) (uint8, error) {
	if c < One || c > Nine {
		return uint8(c), ErrInvalidNumber
	}
	return uint8(c) - Zero, nil
}

func isEvenNumber(i int) bool {
	return i%2 == 0

}

// Puzzle - a sudoku puzzle structure
type Puzzle [9][9]uint8

// Dump - Dump the current state of the puzzle to a writer
func (p *Puzzle) Dump(writer io.Writer) {
	for _, v := range p {
		line := []byte{}
		for i, vv := range v {
			if i != 0 {
				line = append(line, Space)
			}
			if vv == 0 {
				line = append(line, Underscore)
				continue
			}
			line = append(line, vv+Zero)
		}
		// write the line
		line = append(line, Newline)
		writer.Write(line)
	}
}

func puzzleScanSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanLines(data, atEOF)
	if err == nil && token != nil {
		if len(token) != MaxInputRowLength {
			// line length is incorrect, error
			err = ErrInvalidLineLength
			return
		}
		// check that each line is correct format
		for i, b := range token {
			if isEvenNumber(i) {
				// even, should be either a Number or Blank
				if !isNumber(b) && !isBlank(b) {
					//error
					err = ErrInvalidCharacter
					return
				}
			} else {
				// odd, should be space
				if !isSpace(b) {
					err = ErrInvalidCharacter
					return
				}
			}
		}
	}
	return
}

// ParsePuzzle - take an io.Reader and deserialize into a Puzzle
func ParsePuzzle(reader io.Reader) (Puzzle, error) {
	p := Puzzle{}
	// use a scanner to validate, and parse input
	scanner := bufio.NewScanner(reader)
	// use a custom splitter, to break tokens into lines, and validate each line
	// for correctness
	scanner.Split(puzzleScanSplit)
	rowCount := 0
	// scan one row at a tim
	for scanner.Scan() {
		if rowCount > 8 {
			// we have exceeded the allowable number of rows, report invalid
			// row count
			return p, ErrInvalidRowCount
		}
		// grab the token bytes
		token := scanner.Bytes()
		posCount := 0
		for i := 0; i < len(token); i += 2 {
			// since we have already validated the correctness
			// of the puzzle input, we will skip to every other
			// value from the line
			var value uint8 = 0
			if token[i] != Underscore {
				// if the value is not an underscore, set to
				// the number value of the ascii token
				value, _ = asciiToNumber(token[i])
			}
			// populate the value in the matrix
			p[rowCount][posCount] = value
			posCount++
		}
		rowCount++
	}

	if err := scanner.Err(); err != nil {
		// if there are errors, return the errors
		return p, err
	}

	if rowCount < 8 {
		// we have exceeded the allowable number of rows, report invalid
		// row count
		return p, ErrInvalidRowCount
	}

	return p, nil
}

var (
	stuck          = false
	recursionDepth = 30
)

func (p *Puzzle) checkrow(i, j int, k uint8) bool {
	for x := 0; x < 9; x++ {
		if p[i][x] == k {
			return false
		}
	}
	return true
}
func (p *Puzzle) checkcol(i, j int, k uint8) bool {
	for x := 0; x < 9; x++ {
		if p[x][j] == k {
			return false
		}
	}
	return true

}

func (p *Puzzle) checkbox(i, j int, k uint8) bool {

	minX := 3 * int((i)/3)
	minY := 3 * int((j)/3)
	maxX := minX + 3
	maxY := minY + 3

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			if p[x][y] == k {
				return false
			}
		}
	}
	return true
}

func (p *Puzzle) allowed(i, j int, k uint8) bool {
	return p.checkrow(i, j, k) && p.checkcol(i, j, k) && p.checkbox(i, j, k)
}

func (p *Puzzle) isSolved() bool {
	for _, v := range p {
		for _, vv := range v {
			if vv == 0 {
				return false
			}
		}
	}
	return true
}

// BacktrackSolve - solve using backtrack algorithm
func (p *Puzzle) BacktrackSolve() bool {
	for i, _ := range p {
		for j, _ := range p[i] {
			if p[i][j] == 0 {
				// to be filled in
				var k uint8 = 1
				for ; k < 10; k++ {
					if p.allowed(i, j, k) {
						var tmp Puzzle = *p
						tmp[i][j] = k
						if tmp.isSolved() || tmp.BacktrackSolve() {
							*p = tmp
							return true
						}
					}
				}
				return false
			}
		}
	}
	return false
}
