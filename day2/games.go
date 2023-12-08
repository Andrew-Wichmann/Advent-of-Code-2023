package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	red   int
	blue  int
	green int
}

type Game struct {
	id     int
	rounds []Round
}

func parseInput() []Game {
	fd, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	games := []Game{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		gameIdStr := strings.Split(split[0], " ")[1]
		gameId, err := strconv.Atoi(gameIdStr)
		if err != nil {
			panic(err)
		}

		game := Game{id: gameId}
		for _, roundStr := range strings.Split(split[1], "; ") {
			round := Round{}
			for _, color := range strings.Split(roundStr, ", ") {
				colorSplit := strings.Split(color, " ")
				colorValueStr := colorSplit[0]
				colorName := colorSplit[1]

				colorValue, err := strconv.Atoi(colorValueStr)
				if err != nil {
					panic(err)
				}

				switch colorName {
				case "red":
					round.red = colorValue
				case "blue":
					round.blue = colorValue
				case "green":
					round.green = colorValue
				default:
					panic("Invalid color")
				}
			}
			game.rounds = append(game.rounds, round)
		}

		games = append(games, game)
	}
	return games
}
