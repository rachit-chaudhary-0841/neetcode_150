package arrays

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums)/2)

	for _, number := range nums {
		_, isPresent := seen[number]
		if isPresent {
			return true
		}
		seen[number] = struct{}{}
	}

	return false
}
