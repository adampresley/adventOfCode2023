package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adampresley/advent2023/pkg/fileutils"
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
		err           error
		inputLines    []string
		games         []Game
		possibleGames []Game
		gameIDSum     int
	)

	if inputLines, err = fileutils.ReadFileLines("./input.txt"); err != nil {
		fmt.Printf("Error reading input file: %s\n", err.Error())
		os.Exit(-1)
	}

	games = parseInput(inputLines)
	fmt.Printf("Found %d games\n", len(games))
	// fmt.Printf("Game: %+v\n", games)

	for _, game := range games {
		if isGamePossible(game) {
			possibleGames = append(possibleGames, game)
		}
	}

	fmt.Printf("Found %d possible games\n", len(possibleGames))

	for _, game := range possibleGames {
		gameIDSum += game.GameID
	}

	fmt.Printf("Game ID sum: %d\n", gameIDSum)
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

func isGamePossible(game Game) bool {
	for _, round := range game.Rounds {
		if round.Red > MaxRed || round.Green > MaxGreen || round.Blue > MaxBlue {
			return false
		}
	}

	return true
}
