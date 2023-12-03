package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"unicode"

	"github.com/adampresley/advent2023/pkg/fileutils"
)

func main() {
	var (
		err       error
		fileLines []string
		sum       int
	)

	if fileLines, err = fileutils.ReadFileLines("./input.txt"); err != nil {
		slog.Error("error reading input file", "error", err)
		os.Exit(-1)
	}

	for _, line := range fileLines {
		value := getNumberFromLine(line)
		sum += value

		slog.Info(fmt.Sprintf("Line: %s, Value: %d, Sum: %d", line, value, sum))
	}

	slog.Info(fmt.Sprintf("Final sum: %d", sum))
}

func getNumberFromLine(line string) int {
	firstDigit, lastDigit := findFirstAndLastDigits(line)
	s := fmt.Sprintf("%d%d", firstDigit, lastDigit)
	result, _ := strconv.Atoi(s)
	return result
}

func findFirstAndLastDigits(line string) (int, int) {
	var (
		firstDigit int = -1000
		lastDigit  int = -1000
	)

	for _, r := range line {
		if unicode.IsDigit(r) {
			if firstDigit == -1000 {
				firstDigit, _ = strconv.Atoi(string(r))
			} else {
				lastDigit, _ = strconv.Atoi(string(r))
			}
		}
	}

	if lastDigit == -1000 {
		lastDigit = firstDigit
	}

	return firstDigit, lastDigit
}
