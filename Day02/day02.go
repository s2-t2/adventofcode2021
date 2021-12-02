package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func Puzzle1(inputs []string) {
	horizontalPosition := 0
	depth := 0

	for _, input := range inputs {
		splittedInput := strings.Split(input, " ")
		currentCommand := splittedInput[0]
		currentValue, _ := strconv.Atoi(splittedInput[1])
		switch currentCommand {
		case "forward":
			horizontalPosition = horizontalPosition + currentValue
		case "down":
			depth = depth + currentValue
		case "up":
			depth = depth - currentValue
		}
	}
	result := horizontalPosition * depth
	fmt.Println("Day 2 - Puzzle 1:", result)
}

func Puzzle2(inputs []string) {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, input := range inputs {
		splittedInput := strings.Split(input, " ")
		currentCommand := splittedInput[0]
		currentValue, _ := strconv.Atoi(splittedInput[1])
		switch currentCommand {
		case "forward":
			horizontalPosition = horizontalPosition + currentValue
			if aim != 0 {
				depth = depth + (currentValue * aim)
			}
		case "down":
			aim = aim + currentValue
		case "up":
			aim = aim - currentValue
		}
	}
	result := horizontalPosition * depth
	fmt.Println("Day 2 - Puzzle 2:", result)
}
