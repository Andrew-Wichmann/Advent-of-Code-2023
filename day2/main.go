package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an integer argument")
		os.Exit(1)
	}

	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid integer argument")
		os.Exit(1)
	}

	switch arg {
	case 1:
		part1()
	case 2:
		part2()
	default:
		fmt.Println("Invalid argument. Please provide either 1 or 2")
	}
}
