package arrays

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freqMap [26]int16

	for index := range s {
		freqMap[s[index]-'a']++
		freqMap[t[index]-'a']--
	}

	for _, freq := range freqMap {
		if freq != 0 {
			return false
		}
	}

	return true
}
