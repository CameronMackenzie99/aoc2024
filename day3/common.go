package day3

import (
	"os"
	"regexp"
)

func ReadInstructions(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic("ahh")
	}
	instructions := string(content)

	println(instructions)

	return instructions
}

var MulRegex = regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)

func AddUpMatches(matches []string) int {
	result := 0
	re2 := regexp.MustCompile(`[0-9]+`)
	for _, match := range matches {
		digits := re2.FindAllString(match, 2)
		result += convertToInt(digits[0]) * convertToInt(digits[1])
	}
	return result
}
