package day04

import (
	"fmt"
	"strings"
)

type board struct {
	grid [5][5]string
}

func Puzzle1(inputs []string) {
	// Get board
	boards, drawingOrders := GetBoardsAndDrawLists(inputs)
	for _, board := range boards {
		fmt.Println(board)
	}
	fmt.Println(drawingOrders)
	fmt.Println("Day 4 - Puzzle 1:")
}

func GetBoardsAndDrawLists(inputs []string) ([]*board, []string) {
	drawingOrders := strings.Split(inputs[0], ",")
	var inputLines []string
	for i := 2; i < len(inputs); i++ {
		if len(inputs[i]) != 0 {
			inputLines = append(inputLines, inputs[i])
		}
	}

	// initialize 5x5 boards of size totalLen/5
	boards := make([]*board, len(inputLines)/5)
	var tempBoard [5][5]string
	tempCount := 0
	for i := 0; i < len(inputLines); i += 5 {
		//make a temp board to append to actual board ..
		for j, line := range inputLines[i : i+5] {
			line2 := strings.Split(strings.ReplaceAll(strings.TrimLeft(line, " "), " ", " "), " ")
			for _, elem := range line2 {
				tempBoard[tempCount][j] = elem
			}
		}
		tempCount++
		boards[i/5] = &board{grid: tempBoard}
	}

	return boards, drawingOrders
}

func Puzzle2(inputs []string) {
	fmt.Println("Day 4 - Puzzle 2:")
}
