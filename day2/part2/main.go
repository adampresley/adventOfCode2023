package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/adampresley/advent2023/pkg/fileutils"
	"github.com/adampresley/advent2023/pkg/sorting"
)

const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

type Game struct {
	OriginalLine string
	GameID       int
	Rounds       []Round
}

type Round struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	var (
		err         error
		inputLines  []string
		games       []Game
		sumOfPowers int
	)

	if inputLines, err = fileutils.ReadFileLines("./input.txt"); err != nil {
		fmt.Printf("Error reading input file: %s\n", err.Error())
		os.Exit(-1)
	}

	games = parseInput(inputLines)
	fmt.Printf("Found %d games\n", len(games))

	for _, game := range games {
		minimumSet := getMinimumSet(game.Rounds)
		power := minimumSet.Red * minimumSet.Blue * minimumSet.Green
		sumOfPowers += power
	}

	fmt.Printf("Sum of powers: %d\n", sumOfPowers)
}

func parseInput(inputLines []string) []Game {
	var (
		err   error
		games []Game
	)

	for lineNumber, line := range inputLines {
		var (
			game Game
		)

		if game, err = parseGame(line); err != nil {
			fmt.Printf("Error parsing game on line  number %d: %s\n", lineNumber, err.Error())
			os.Exit(-1)
		}

		games = append(games, game)
	}

	return games
}

func parseGame(line string) (Game, error) {
	var (
		err  error
		game Game
	)

	game.OriginalLine = line
	idAndRoundsSplit := strings.Split(line, ": ")

	if len(idAndRoundsSplit) < 2 {
		return game, fmt.Errorf("Invalid game format on line: %s", line)
	}

	game.GameID = getGameID(idAndRoundsSplit[0])

	if game.Rounds, err = parseRounds(idAndRoundsSplit[1]); err != nil {
		return game, err
	}

	return game, err
}

func getGameID(part string) int {
	s := strings.Split(part, " ")

	if len(s) < 2 {
		return 0
	}

	id, _ := strconv.Atoi(s[1])
	return id
}

func parseRounds(rounds string) ([]Round, error) {
	var (
		err    error
		result []Round
		round  Round
	)

	roundsSplit := strings.Split(rounds, "; ")

	for _, r := range roundsSplit {
		if round, err = parseRound(r); err != nil {
			return result, err
		}

		result = append(result, round)
	}

	return result, err
}

func parseRound(round string) (Round, error) {
	var (
		err    error
		result Round
	)

	roundSplit := strings.Split(round, ", ")

	for _, r := range roundSplit {
		numAndColorSpit := strings.Split(r, " ")

		if len(numAndColorSpit) < 2 {
			return result, fmt.Errorf("Invalid round format: %s", round)
		}

		num, _ := strconv.Atoi(numAndColorSpit[0])
		color := numAndColorSpit[1]

		switch color {
		case "red":
			result.Red = num

		case "green":
			result.Green = num

		case "blue":
			result.Blue = num
		}
	}

	return result, err
}

func getMinimumSet(rounds []Round) Round {
	var (
		reds   []int
		greens []int
		blues  []int
	)

	for _, round := range rounds {
		reds = append(reds, round.Red)
		greens = append(greens, round.Green)
		blues = append(blues, round.Blue)
	}

	slices.SortFunc(reds, sorting.IntReverse)
	slices.SortFunc(greens, sorting.IntReverse)
	slices.SortFunc(blues, sorting.IntReverse)

	return Round{
		Red:   reds[0],
		Green: greens[0],
		Blue:  blues[0],
	}
}
