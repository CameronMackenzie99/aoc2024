package day1

func calculateOccurences(list1 []int) map[int]int {
	occurences1 := make(map[int]int)

	for i := 0; i < len(list1); i++ {
		occurences1[list1[i]] += 1
	}

	return occurences1

}

func Part2() int {
	list1, list2 := ReadLists("day1/input.txt")

	occurences1 := calculateOccurences(list1)
	occurences2 := calculateOccurences(list2)

	similarityScore := 0

	for id, occ1 := range occurences1 {
		if occ2, ok := occurences2[id]; ok {
			similarityScore += occ2 * occ1 * id
		}
	}

	// fmt.Printf("%v", occurences1)
	// fmt.Printf("%v", occurences2)

	return similarityScore
}
