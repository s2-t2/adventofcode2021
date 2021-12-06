package main

import (
	"bufio"
	"log"
	"os"

	day06 "github.com/sujata5538/adventofcode2021/Day06"
)

func main() {
	// Day 1:
	// dayOneInput := ReadFromFile("data/day01.txt")
	// day01.Puzzle1(dayOneInput)
	// day01.Puzzle2(dayOneInput)

	// // // Day 2:
	// dayTwoInput := ReadFromFile("data/day02.txt")
	// day02.Puzzle1(dayTwoInput)
	// day02.Puzzle2(dayTwoInput)

	// // Day 3:
	// dayThreeInput := ReadFromFile("data/day03.txt")
	// day03.Puzzle1(dayThreeInput)
	// day03.Puzzle2(dayThreeInput)

	// Day 3:
	// dayFourInput := ReadFromFile("data/day04.txt")
	// day04.Puzzle1(dayFourInput)
	// day04.Puzzle2(dayFourInput)

	// Day 6:
	daySixInput := ReadFromFile("data/day06.txt")
	// day06.Puzzle1(daySixInput)
	day06.Puzzle2(daySixInput)
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
