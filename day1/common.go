package day1

import (
	"os"
	"strconv"
	"strings"
)

func ReadLists(file string) (list1, list2 []int) {
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
