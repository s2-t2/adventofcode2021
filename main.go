package main

import (
	"bufio"
	"log"
	"os"

	day01 "github.com/sujata5538/adventofcode2021/Day01"
	day02 "github.com/sujata5538/adventofcode2021/Day02"
	day03 "github.com/sujata5538/adventofcode2021/Day03"
)

func main() {
	// Day 1:
	dayOneInput := ReadFromFile("data/day01.txt")
	day01.Puzzle1(dayOneInput)
	day01.Puzzle2(dayOneInput)

	// // Day 2:
	dayTwoInput := ReadFromFile("data/day02.txt")
	day02.Puzzle1(dayTwoInput)
	day02.Puzzle2(dayTwoInput)

	// Day 3:
	dayThreeInput := ReadFromFile("data/day03.txt")
	day03.Puzzle1(dayThreeInput)
	day03.Puzzle2(dayThreeInput)
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
