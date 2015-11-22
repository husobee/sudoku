package sudoku

import "bufio"

const (
	underscore        byte = 0x5f
	space                  = 0x20
	zero                   = 0x30
	one                    = 0x31
	nine                   = 0x39
	newline                = 0x0a
	maxInputRowLength int  = 17
)

func isSpace(c byte) bool {
	return c == space
}

func isBlank(c byte) bool {
	return c == underscore
}

func isNumber(c byte) bool {
	if _, err := asciiToNumber(c); err != nil {
		return false
	}
	return true
}

func asciiToNumber(c byte) (uint8, error) {
	if c < one || c > nine {
		return uint8(c), ErrParseInvalidNumber
	}
	return uint8(c) - zero, nil
}

func isEvenNumber(i int) bool {
	return i%2 == 0
}

func puzzleScanSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	advance, token, err = bufio.ScanLines(data, atEOF)
	if err == nil && token != nil {
		if len(token) != maxInputRowLength {
			// line length is incorrect, error
			err = ErrParseInvalidLineLength
			return
		}
		// check that each line is correct format
		for i, b := range token {
			if isEvenNumber(i) {
				// even, should be either a Number or Blank
				if !isNumber(b) && !isBlank(b) {
					//error
					err = ErrParseInvalidCharacter
					return
				}
			} else {
				// odd, should be space
				if !isSpace(b) {
					err = ErrParseInvalidCharacter
					return
				}
			}
		}
	}
	return
}
