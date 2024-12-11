package day3

import (
	"os"
	"regexp"
	"strings"
)

func readMultiLineInstructions(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		panic("ahh")
	}
	instructions := strings.Join(strings.Split(string(content), "\n"), "")

	println(instructions)

	return instructions
}

func excludeDisabledInstructions(instructions string) string {
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don\'t\(\)`)

	lastDoIndexRanges := doRegex.FindAllStringIndex(instructions, -1)
	lastDontIndexRanges := dontRegex.FindAllStringIndex(instructions, -1)

	lastDo, lastDont := lastDoIndexRanges[len(lastDoIndexRanges)-1], lastDontIndexRanges[len(lastDontIndexRanges)-1]

	if lastDo[0] < lastDont[0] {
		return instructions[:lastDont[0]]
	}

	instructionsArray := strings.Split(instructions, "")

	return strings.Join(append(instructionsArray[:lastDont[0]], instructionsArray[lastDo[0]:]...), "")
}

func Part2() int {
	validInstructions := ReadInstructions("day3/input.txt")
	// validInstructions := excludeDisabledInstructions(instructions)
	matches := MulRegex.FindAllString(validInstructions, -1)

	return AddUpMatches(matches)
}
