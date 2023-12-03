package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/adampresley/advent2023/pkg/fileutils"
)

const (
	debug = false

	STATE_PARSING = iota
	STATE_IN_NUMBER
)

type Line struct {
	line         string
	previousLine *Line
	nextLine     *Line
	tokens       []Token
}

type Token struct {
	character string
	pos       int
	isDigit   bool
	isSymbol  bool
}

func main() {
	var (
		err   error
		lines []string
		sum   int
	)

	filename := "input.txt"

	if debug {
		filename = "sample.txt"
	}

	if lines, err = fileutils.ReadFileLines(filename); err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		os.Exit(-1)
	}

	/*
	 * Tokenize the input
	 */
	tokenizedLines := tokenizeLines(lines)
	matrix := buildMatrix(tokenizedLines)

	for _, line := range matrix {
		fmt.Printf("%s", line)
		sum += parseLine(line)
	}

	fmt.Printf("\nSum: %d\n", sum)
}

func tokenizeLines(lines []string) []*Line {
	var (
		result []*Line
	)

	for _, line := range lines {
		result = append(result, &Line{
			line:   line,
			tokens: tokenizeLine(line),
		})
	}

	return result
}

func tokenizeLine(line string) []Token {
	var (
		tokens []Token
	)

	for index, character := range line {
		token := Token{
			character: string(character),
			pos:       index,
			isDigit:   false,
			isSymbol:  false,
		}

		if unicode.IsDigit(character) {
			token.isDigit = true
		}

		if isSymbol(character) {
			token.isSymbol = true
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func buildMatrix(lines []*Line) []*Line {
	var (
		result []*Line
	)

	for index, line := range lines {
		result = append(result, line)

		if index > 0 {
			result[index].previousLine = lines[index-1]
		}

		if index < len(lines)-2 {
			result[index].nextLine = lines[index+1]
		}
	}

	return result
}

/*
parseLine parses a line of tokens, checking for numbers that have adjacent
symbols. This returns a sum of all numbers that have adjacent symbols.
*/
func parseLine(line *Line) int {
	var (
		state                  int = STATE_PARSING
		capturedNumberString   string
		capturedNumberStartPos int
		capturedNumber         int
		sum                    int
	)

	for _, token := range line.tokens {
		if token.isDigit {
			if state == STATE_PARSING {
				state = STATE_IN_NUMBER
				capturedNumberStartPos = token.pos
			}

			capturedNumberString += token.character
		}

		if !token.isDigit && state == STATE_IN_NUMBER {
			capturedNumber = captureNumber(capturedNumberString, capturedNumberStartPos, line)
			sum += capturedNumber

			capturedNumberString = ""
			capturedNumberStartPos = 0
		}

		if !token.isDigit {
			state = STATE_PARSING
		}
	}

	if state == STATE_IN_NUMBER {
		capturedNumber = captureNumber(capturedNumberString, capturedNumberStartPos, line)
		sum += capturedNumber
		capturedNumberString = ""
	}

	return sum
}

func captureNumber(s string, startPos int, line *Line) int {
	var (
		result   int
		hasAbove bool
		hasLeft  bool
		hasBelow bool
		hasRight bool
		startCol int
		endCol   int
		p        *Line
		n        *Line
	)

	/*
	 * Find the start and end columns to search for symbols on
	 */
	startCol = startPos - 1
	endCol = startPos + len(s)

	if startCol < 0 {
		startCol = 0
	}

	if endCol > len(line.line)-1 {
		endCol = len(line.line) - 1
	}

	/*
	 * Check the previous line for adjacent symbols
	 */
	if line.previousLine != nil {
		p = line.previousLine

		for index := startCol; index <= endCol; index++ {
			if p.tokens[index].isSymbol {
				hasAbove = true
			}
		}
	}

	/*
	 * Check the next line for adjacent symbols
	 */
	if line.nextLine != nil {
		n = line.nextLine

		for index := startCol; index <= endCol; index++ {
			if n.tokens[index].isSymbol {
				hasBelow = true
			}
		}
	}

	/*
	 * Check left and right of number for adjacent symbols
	 */
	if line.tokens[startCol].isSymbol {
		hasLeft = true
	}

	if line.tokens[endCol].isSymbol {
		hasRight = true
	}

	if hasAbove || hasBelow || hasLeft || hasRight {
		result, _ = strconv.Atoi(s)
		fmt.Printf("  Captured number %d at pos %d because of adjacent symbols\n", result, startPos)
	}

	return result
}

func isSymbol(character rune) bool {
	return !unicode.IsDigit(character) && character != '.'
}

func (l Line) String() string {
	result := "Line: " + l.line + "\n"
	return result
}

func (t Token) String() string {
	return fmt.Sprintf("  '%s': isDigit: %v, isSymbol: %v", t.character, t.isDigit, t.isSymbol)
}
