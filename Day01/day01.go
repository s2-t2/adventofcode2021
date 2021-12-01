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
		A, _ := strconv.Atoi(inputs[i])
		B, _ := strconv.Atoi(inputs[i+1])
		C, _ := strconv.Atoi(inputs[i+2])
		D, _ := strconv.Atoi(inputs[i+3])
		firstSum := A + B + C
		secondSum := B + C + D
		if secondSum > firstSum {
			totalIncreased++
		}
	}
	fmt.Println("Day 1 - Puzzle 2:", totalIncreased)
}
