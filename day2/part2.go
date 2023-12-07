package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ok := scanner.Scan()
	if !ok {
		panic("Error reading file")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	times := strings.Fields(scanner.Text())[1:]
	ok = scanner.Scan()
	if !ok {
		panic("Error reading file")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	distances := strings.Fields(scanner.Text())[1:]

	joinedTimes, err := strconv.Atoi(strings.Join(times, ""))
	if err != nil {
		panic(err)
	}
	joinedDistances, err := strconv.Atoi(strings.Join(distances, ""))
	if err != nil {
		panic(err)
	}

	races := []Race{Race{Time: joinedTimes, Distance: joinedDistances}}
	solve(races)
}
