package main

import "math"

func solveQuadratic(b, c float64) (float64, float64) {
	// calculate the discriminant
	d := b*b - 4*c

	// calculate the two roots
	root1 := (-b + math.Sqrt(d)) / -2
	root2 := (-b - math.Sqrt(d)) / -2

	return root1, root2
}

func product(nums []int) int {
	result := 1
	for _, num := range nums {
		result *= num
	}
	return result
}

func solve(races []Race) error {
	solutions := []int{}
	for _, race := range races {
		min, max := solveQuadratic(float64(race.Time), float64(race.Distance))
		solutions = append(solutions, (int(math.Floor(max) - math.Ceil(min) + 1)))
	}

	println(product(solutions))
	return nil
}
