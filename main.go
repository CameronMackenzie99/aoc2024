package main

import (
	"aoc2024/day2"
	"fmt"
	"time"
)

func measureAndPrint[T any](label string, fn func() T) {
	startTime := time.Now()
	res := fn()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("%s: %-15v\t(%v)\n", label, res, elapsed)
}

func main() {
	// measureAndPrint("Day1Part1", day1.Part1)
	// measureAndPrint("Day1Part2", day1.Part2)
	// measureAndPrint("Day2Part1", day2.Part1)
	measureAndPrint("Day2Part2", day2.Part2)

}
