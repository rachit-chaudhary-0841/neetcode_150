package arrays

func groupAnagrams(strs []string) [][]string {
	groupIndexMap := make(map[[26]uint8]int, len(strs)/2)
	groups := make([][]string, 0, len(strs)/2)

	for _, str := range strs {
		var hash [26]uint8

		for index := 0; index < len(str); index++ {
			hash[str[index]-'a']++
		}

		groupIndex, isPresent := groupIndexMap[hash]
		if isPresent {
			groups[groupIndex] = append(groups[groupIndex], str)
		} else {
			groupIndexMap[hash] = len(groupIndexMap)
			groups = append(groups, []string{str})
		}
	}

	return groups
}
