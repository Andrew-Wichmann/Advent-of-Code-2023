package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Part1() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var runningTotal int
	var first, second rune
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		// forwards
		for _, char := range line {
			if unicode.IsDigit(char) {
				first = char
				break
			}
		}
		// backwards
		for j := len(line) - 1; j >= 0; j-- {
			char := runes[j]
			if unicode.IsDigit(char) {
				second = char
				break
			}
		}
		number := string(first) + string(second)

		// Convert the number to an integer
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Error converting number:", err)
			return
		}

		// Add the number to a running total
		runningTotal += num

	}
	fmt.Println(runningTotal)

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}
