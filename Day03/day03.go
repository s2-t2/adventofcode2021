package day03

import (
	"fmt"
	"strconv"
	"strings"
)

func Puzzle1(inputs []string) {
	// Find the most common bit for each bits
	for _, v := range inputs {
		individualChars := strings.Split(v, "")
		// Get most common bit for individual chars, only first bit, then, only second bit etc....
		fmt.Println(individualChars)
	}

	// Refactor the whole thing ...
	var firstBits []string
	var secondBits []string
	var thirdBits []string
	var fourthBits []string
	var fifthBits []string
	var sixthBits []string
	var seventhBits []string
	var eighthBits []string
	var ninthBits []string
	var tenthBits []string
	var eleventhBits []string
	var twelthBits []string

	for _, value := range inputs {
		characters := strings.Split(value, "")
		firstBits = append(firstBits, characters[0])
		secondBits = append(secondBits, characters[1])
		thirdBits = append(thirdBits, characters[2])
		fourthBits = append(fourthBits, characters[3])
		fifthBits = append(fifthBits, characters[4])
		sixthBits = append(sixthBits, characters[5])
		seventhBits = append(seventhBits, characters[6])
		eighthBits = append(eighthBits, characters[7])
		ninthBits = append(ninthBits, characters[8])
		tenthBits = append(tenthBits, characters[9])
		eleventhBits = append(eleventhBits, characters[10])
		twelthBits = append(twelthBits, characters[11])
	}

	var result string
	firstBitResult := GetMostSignificantBit(firstBits)
	secondBitResult := GetMostSignificantBit(secondBits)
	thirdBitResult := GetMostSignificantBit(thirdBits)
	fourthBitResult := GetMostSignificantBit(fourthBits)
	fifthBitsResult := GetMostSignificantBit(fifthBits)
	sixthBitsResult := GetMostSignificantBit(sixthBits)
	seventhBitsResult := GetMostSignificantBit(seventhBits)
	eighthBitsResult := GetMostSignificantBit(eighthBits)
	ninthBitsResult := GetMostSignificantBit(ninthBits)
	tenthBitsResult := GetMostSignificantBit(tenthBits)
	eleventhBitsResult := GetMostSignificantBit(eleventhBits)
	twelthBitsResult := GetMostSignificantBit(twelthBits)
	result += firstBitResult
	result += secondBitResult
	result += thirdBitResult
	result += fourthBitResult
	result += fifthBitsResult
	result += sixthBitsResult
	result += seventhBitsResult
	result += eighthBitsResult
	result += ninthBitsResult
	result += tenthBitsResult
	result += eleventhBitsResult
	result += twelthBitsResult

	// Flip bits
	var leastSignificantResult string
	individualResults := strings.Split(result, "")
	for _, value := range individualResults {
		if value == "0" {
			leastSignificantResult += "1"
		} else {
			leastSignificantResult += "0"
		}
	}

	// fmt.Println(result)
	// fmt.Println(leastSignificantResult)

	// Convert string to binary to decimal
	var firstRes int64
	var secondRes int64
	if i, err := strconv.ParseInt(result, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		firstRes = i
	}

	if i, err := strconv.ParseInt(leastSignificantResult, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		secondRes = i
	}

	fmt.Println("Day 3 - Puzzle 1:", firstRes*secondRes)
}

// Helper function
func GetMostSignificantBit(inputBits []string) string {
	var result string
	numberOf0s := 0
	numberOf1s := 0
	for _, bit := range inputBits {
		if bit == "0" {
			numberOf0s++
		}
		if bit == "1" {
			numberOf1s++
		}
	}
	if numberOf0s > numberOf1s {
		result += "0"
	} else {
		result += "1"
	}

	return result
}

func Puzzle2(inputs []string) {
	var updatedMostBitsToConsider = inputs
	for i := 0; i < 12; i++ {
		significantBitAs1List, significantBitAs0List := GetSignificantBitsList(updatedMostBitsToConsider, i)
		if len(significantBitAs1List) > len(significantBitAs0List) {
			updatedMostBitsToConsider = significantBitAs1List
		} else if len(significantBitAs0List) > len(significantBitAs1List) {
			updatedMostBitsToConsider = significantBitAs0List
		} else if len(significantBitAs0List) == len(significantBitAs1List) {
			updatedMostBitsToConsider = significantBitAs1List
		}
		if len(updatedMostBitsToConsider) <= 0 {
			break
		}
	}

	var oxygenGeneratorRating int64
	if decimalValue, err := strconv.ParseInt(updatedMostBitsToConsider[0], 2, 64); err != nil {
		fmt.Println(err)
	} else {
		oxygenGeneratorRating = decimalValue
	}

	// C02 scrubber
	updatedLeastBitsToConsider := inputs
	for i := 0; i < 12; i++ {
		significantBitAs1List, significantBitAs0List := GetSignificantBitsList(updatedLeastBitsToConsider, i)
		if len(significantBitAs1List) < len(significantBitAs0List) {
			updatedLeastBitsToConsider = significantBitAs1List
		} else if len(significantBitAs0List) < len(significantBitAs1List) {
			updatedLeastBitsToConsider = significantBitAs0List
		} else if len(significantBitAs0List) == len(significantBitAs1List) {
			updatedLeastBitsToConsider = significantBitAs0List
		}
		if len(updatedLeastBitsToConsider) <= 1 {
			break
		}
	}

	var co2ScrubberRating int64
	if decimalVaue, err := strconv.ParseInt(updatedLeastBitsToConsider[0], 2, 64); err != nil {
		fmt.Println(err)
	} else {
		co2ScrubberRating = decimalVaue
	}

	lifeSupport := oxygenGeneratorRating * co2ScrubberRating
	fmt.Println("Day 3 - Puzzle 2:", lifeSupport)
}

// Helper function for Puzzle 2
func GetSignificantBitsList(listToConsider []string, indexToTake int) ([]string, []string) {
	var bitsWith1AsSignificantBit []string
	var bitsWith0AsSignificantBit []string
	for _, value := range listToConsider {
		individualChars := strings.Split(value, "")
		if individualChars[indexToTake] == "1" {
			bitsWith1AsSignificantBit = append(bitsWith1AsSignificantBit, value)
		} else if individualChars[indexToTake] == "0" {
			bitsWith0AsSignificantBit = append(bitsWith0AsSignificantBit, value)
		}
	}
	return bitsWith1AsSignificantBit, bitsWith0AsSignificantBit
}
