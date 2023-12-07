package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
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
	races := []Race{}
	for i := 0; i < len(times); i++ {
		// Convert the number to an integer
		time, err := strconv.Atoi(times[i])
		if err != nil {
			fmt.Println("Error converting number:", err)
			return
		}
		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			fmt.Println("Error converting number:", err)
			return
		}

		races = append(races, Race{Time: time, Distance: distance})
	}
	solve(races)
}
