package main

import (
	"bufio"
	"log"
	"os"

	day01 "github.com/sujata5538/adventofcode2021/Day01"
)

func main() {
	// Day 1:
	dayOneInput := ReadFromFile("data/day01.txt")
	day01.Puzzle1(dayOneInput)
	day01.Puzzle2(dayOneInput)
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
