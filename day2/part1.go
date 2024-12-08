package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLevels(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic("ahh")
	}
	lines := strings.Split(strings.TrimSuffix(string(content), "\n"), "\n")
	return lines
}

func checkLevelSafe(level []string, dampener bool) bool {
	netChange := 0
	absChange := 0

	for i := 0; i < len(level)-1; i++ {
		diff := stringToInt(level[i]) - stringToInt(level[i+1])
		absDiff := absDiffInt(stringToInt(level[i]), stringToInt(level[i+1]))
		netChange += diff
		absChange += absDiff
		// println("level")
		if absDiff > 3 || absDiff == 0 || absInt(netChange) != absChange {
			fmt.Println("level")
			fmt.Printf("%v", level)
			fmt.Println()
			if dampener {
				shortenedLevel := make([]string, len(level))
				copy(shortenedLevel, level)
				shortenedLevel = append(shortenedLevel[:i], shortenedLevel[i+1:]...)
				fmt.Println("shortenedLevel")
				fmt.Printf("%v", shortenedLevel)
				fmt.Println()

				if i == len(level)-2 {
					shortenedLevel := make([]string, len(level))
					copy(shortenedLevel, level)
					shortenedLevel = shortenedLevel[:len(shortenedLevel)-1]
					fmt.Println("shortenedLevel")
					fmt.Printf("%v", shortenedLevel)
					fmt.Println()
					if checkLevelSafe(shortenedLevel, false) {
						return true
					}
				}
				return checkLevelSafe(shortenedLevel, false)
			}
			return false

		}

	}

	return true
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func stringToInt(valString string) int {
	val, err := strconv.Atoi(valString)
	if err != nil {
		panic(err)
	}
	return val
}

func Part1() int {
	levels := readLevels("day2/input.txt")

	safeCount := 0
	for _, level := range levels {
		sequence := strings.Split(level, " ")
		if checkLevelSafe(sequence, false) {
			safeCount += 1
		}
	}
	return safeCount
}
