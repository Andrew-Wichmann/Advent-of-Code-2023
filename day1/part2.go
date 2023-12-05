package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var MAPPING = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// mapStringsToNumbers converts a string into a slice of integers.
func mapStringsToNumbers(line string) []string {
	var numbers []string
	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			numbers = append(numbers, string(line[i]))
		}
		for j := i + 1; j <= len(line); j++ {
			if val, ok := MAPPING[line[i:j]]; ok {
				numbers = append(numbers, val)
			}
		}
	}
	return numbers
}

func Part2() {
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
	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()

		numbers := mapStringsToNumbers(line)

		// Concatenate the first and last numbers and convert to an integer
		first := numbers[0]
		last := numbers[len(numbers)-1]
		num, err := strconv.Atoi(first + last)
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
