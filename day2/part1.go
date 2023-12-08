package main

const redMax = 12
const blueMax = 14
const greenMax = 13

func part1() {
	sum := 0
	games := parseInput()
	for _, game := range games {
		possible := true
		for _, round := range game.rounds {
			if round.red > redMax || round.blue > blueMax || round.green > greenMax {
				possible = false
			}
		}
		if possible {
			sum += game.id
		}
	}
	println(sum)
}
