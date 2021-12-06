package day06

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func Puzzle1(inputs []string) {
	inputFishTimer := strings.Split(inputs[0], ",")
	currentFishTimer := inputFishTimer
	tempFishTimer := inputFishTimer

	for i := 0; i < 80; i++ {
		for i := 0; i < len(currentFishTimer); i++ {
			if currentFishTimer[i] == "0" {
				tempFishTimer[i] = "6"
				tempFishTimer = append(tempFishTimer, "8")
				currentFishTimer[i] = "6"
			} else {
				currentValue, _ := strconv.Atoi(currentFishTimer[i])
				tempFishTimer[i] = strconv.Itoa(currentValue - 1)
			}
		}
		currentFishTimer = tempFishTimer
	}
	totalFish := len(currentFishTimer)
	fmt.Println(totalFish)
}

func Puzzle2(inputs []string) {
	inputFishTimer := strings.Split(inputs[0], ",")
	inputFishTimerInt := ConvertStringToInt(inputFishTimer)

	_, length := GetFishTimer(inputFishTimerInt, 150)

	fmt.Println(" current fish timer 1:", length)

}

func GetFishTimer(inputFishTimerInt []big.Int, nrOfIteration int) ([]big.Int, big.Int) {
	var currentFishTimer []big.Int
	currentFishTimer = inputFishTimerInt
	tempFishTimer := inputFishTimerInt
	for i := 0; i < nrOfIteration; i++ {
		for i := 0; i < len(currentFishTimer); i++ {
			if currentFishTimer[i].Cmp(big.NewInt(int64(0))) == 0 {
				tempFishTimer[i] = *big.NewInt(int64(6))
				tempFishTimer = append(tempFishTimer, *big.NewInt(int64(8)))
				currentFishTimer[i] = *big.NewInt(int64(6))
			} else {
				currentValue := currentFishTimer[i]
				tempFishTimer[i] = *currentValue.Sub(&currentValue, big.NewInt(int64(1)))
			}
		}

		currentFishTimer = tempFishTimer
	}

	return currentFishTimer, *big.NewInt(int64(len(currentFishTimer)))
}

func ConvertStringToInt(stringArray []string) []big.Int {
	var result []big.Int
	for _, v := range stringArray {
		var bignum, _ = new(big.Int).SetString(v, 0)
		result = append(result, *bignum)
	}

	return result
}
