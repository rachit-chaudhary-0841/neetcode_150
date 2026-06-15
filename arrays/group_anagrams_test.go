package arrays

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:  "basic anagrams",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name:     "single word",
			input:    []string{"hello"},
			expected: [][]string{{"hello"}},
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: [][]string{},
		},
		{
			name:  "duplicates",
			input: []string{"a", "a", "a"},
			expected: [][]string{
				{"a", "a", "a"},
			},
		},
		{
			name:  "mixed lengths",
			input: []string{"abc", "bca", "cab", "abcd", "bcad"},
			expected: [][]string{
				{"abc", "bca", "cab"},
				{"abcd", "bcad"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.input)

			if !isEqualGroups(result, tt.expected) {
				t.Errorf("got %v, expected %v", result, tt.expected)
			}
		})
	}
}

// isEqualGroups compares two [][]string ignoring order of groups and order within groups
func isEqualGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	used := make([]bool, len(b))

	for _, ga := range a {
		found := false

		for j, gb := range b {
			if used[j] {
				continue
			}
			if isSameGroup(ga, gb) {
				used[j] = true
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

// isSameGroup checks if two slices contain the same elements (ignoring order)
func isSameGroup(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[string]int)

	for _, s := range a {
		counts[s]++
	}

	for _, s := range b {
		if counts[s] == 0 {
			return false
		}
		counts[s]--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}
