package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 3   4
// 4   3
// 2   5
// 1   3
// 3   9
// 3   3

func readLists(file string) (list1, list2 []int) {
	content, err := os.ReadFile(file)
	if err != nil {
		panic("ahh")
	}
	lines := strings.Split(strings.TrimSuffix(string(content), "\n"), "\n")

	list1 = make([]int, len(lines))
	list2 = make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		listEntries := strings.Split(lines[i], "   ")
		entry1, err := strconv.Atoi(listEntries[0])
		if err != nil {
			panic(err)
		}
		entry2, err := strconv.Atoi(listEntries[1])
		if err != nil {
			panic(err)
		}
		list1[i], list2[i] = entry1, entry2
	}

	return list1, list2
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Part1() {
	list1, list2 := readLists("day1/input.txt")

	sortedList1 := mergeSort(list1)
	sortedList2 := mergeSort(list2)

	fmt.Println("list1 len", len(list1))
	fmt.Println("sortedList1 len", len(sortedList1))
	fmt.Println("list2 len", len(list2))
	fmt.Println("sortedList2 len", len(sortedList2))

	totalLength := 0
	for i := 0; i < len(list1); i++ {
		totalLength += absDiffInt(sortedList1[i], sortedList2[i])
	}

	println("totalLength", totalLength)

}
