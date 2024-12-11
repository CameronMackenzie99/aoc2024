package day3

import (
	"strconv"
)

func convertToInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic("ahh")
	}
	return intVal
}

func Part1() int {

	instructions := ReadInstructions("day3/input_1.txt")

	matches := MulRegex.FindAllString(instructions, -1)

	return AddUpMatches(matches)
}
