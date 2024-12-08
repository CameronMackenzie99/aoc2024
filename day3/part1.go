package day3

import (
	"os"
	"regexp"
	"strconv"
)

func convertToInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic("ahh")
	}
	return intVal
}

func readInstructions(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic("ahh")
	}
	instructions := string(content)

	println(instructions)

	return instructions
}

func Part1() int {

	instructions := readInstructions("day3/input_1.txt")

	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)

	matches := re.FindAllString(instructions, -1)

	result := 0
	re2 := regexp.MustCompile(`[0-9]+`)
	for _, match := range matches {
		digits := re2.FindAllString(match, 2)
		result += convertToInt(digits[0]) * convertToInt(digits[1])
	}
	return result
}
