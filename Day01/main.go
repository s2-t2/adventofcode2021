package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputs := ReadFromFile("../data/day01.txt")
	//Part 1
	totalIncreased := 0
	for i := 0; i < len(inputs)-1; i++ {
		firstInput, _ := strconv.Atoi(inputs[i])
		secondInput, _ := strconv.Atoi(inputs[i+1])
		if secondInput > firstInput {
			totalIncreased++
		}
	}
	fmt.Println("Part1 Puzzle result:", totalIncreased)
	// Part 2
	totalIncreased2 := 0
	for i := 0; i < len(inputs)-3; i++ {
		A, _ := strconv.Atoi(inputs[i])
		B, _ := strconv.Atoi(inputs[i+1])
		C, _ := strconv.Atoi(inputs[i+2])
		D, _ := strconv.Atoi(inputs[i+3])
		firstSum := A + B + C
		secondSum := B + C + D
		if secondSum > firstSum {
			totalIncreased2++
		}
	}
	fmt.Println("Part2 Puzzle Result:", totalIncreased2)
}
func ReadFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	return txtlines
}
