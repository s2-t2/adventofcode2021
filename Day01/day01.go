package day01

import (
	"fmt"
	"strconv"
)

func Puzzle1(inputs []string) {
	totalIncreased := 0
	for i := 0; i < len(inputs)-1; i++ {
		firstInput, _ := strconv.Atoi(inputs[i])
		secondInput, _ := strconv.Atoi(inputs[i+1])
		if secondInput > firstInput {
			totalIncreased++
		}
	}
	fmt.Println("Day 1 - Puzzle 1:", totalIncreased)
}

func Puzzle2(inputs []string) {
	totalIncreased := 0
	for i := 0; i < len(inputs)-3; i++ {
		// Checking B+C+D > A+B+C is same as checking D > A
		firstInput, _ := strconv.Atoi(inputs[i])
		fourthInput, _ := strconv.Atoi(inputs[i+3])

		if fourthInput > firstInput {
			totalIncreased++
		}
	}
	fmt.Println("Day 1 - Puzzle 2:", totalIncreased)
}
