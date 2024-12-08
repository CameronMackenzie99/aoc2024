package day2

import (
	"fmt"
	"strings"
)

func Part2() int {
	levels := readLevels("day2/test_input.txt")

	safeCount := 0
	for _, level := range levels {
		sequence := strings.Split(level, " ")
		if checkLevelSafe(sequence, true) {
			println("level safe")
			fmt.Printf("%v", sequence)
			println()
			safeCount += 1
		}
	}
	return safeCount
}
