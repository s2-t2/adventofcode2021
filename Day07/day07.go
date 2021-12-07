package day07

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Puzzle1(inputs []string) {
	input := ConvertStringToInt(strings.Split(inputs[0], ","))
	minPathCost := math.MaxInt
	for i := 0; i < len(input); i++ {
		currentPathCost := FindFuelCostOfMinPath(input, input[i])
		if currentPathCost < minPathCost {
			minPathCost = currentPathCost
		}
	}

	fmt.Println("Day 07 Puzzle 1", minPathCost)
}

func Puzzle2(inputs []string) {
	input := ConvertStringToInt(strings.Split(inputs[0], ","))

	minPathCost := math.MaxInt
	changedByStep := 0
	for i := 0; i < len(input); i++ {
		changedByStep++
		currentPathCost := FindCrabPath(input, changedByStep)
		if currentPathCost < minPathCost {
			minPathCost = currentPathCost
		}
	}

	fmt.Println("Day 07 Puzzle 2", minPathCost)
}

func FindCrabPath(crabPath []int, changeByStep int) int {
	var fuelCost int
	for i := 0; i < len(crabPath); i++ {
		var currentCost int
		if crabPath[i] > changeByStep {
			difference := crabPath[i] - changeByStep
			currentCost = (difference * (difference + 1)) / 2
		} else {
			difference := changeByStep - crabPath[i]
			currentCost = (difference * (difference + 1)) / 2
		}
		fuelCost = fuelCost + currentCost
	}

	return fuelCost
}

func FindFuelCostOfMinPath(minPath []int, changeByStep int) int {
	var fuelCost int
	fuelCost = 0
	for i := 0; i < len(minPath); i++ {
		var currentCost int
		if minPath[i] > changeByStep {
			currentCost = minPath[i] - changeByStep
		} else {
			currentCost = changeByStep - minPath[i]
		}

		fuelCost = fuelCost + currentCost
	}

	return fuelCost
}

func ConvertStringToInt(stringArray []string) []int {
	var result []int
	for _, v := range stringArray {
		intconv, _ := strconv.Atoi(v)
		result = append(result, intconv)
	}

	return result
}
