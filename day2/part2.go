package main

func part2() {
	games := parseInput()
	sum := 0
	for _, game := range games {

		minBlue := 0
		minGreen := 0
		minRed := 0

		for _, round := range game.rounds {
			if round.red > minRed {
				minRed = round.red
			}
			if round.blue > minBlue {
				minBlue = round.blue
			}
			if round.green > minGreen {
				minGreen = round.green
			}
		}
		sum += minRed * minBlue * minGreen
	}
	println(sum)
}
