package arrays

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "has duplicates",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: false,
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: false,
		},
		{
			name:     "all elements same",
			input:    []int{7, 7, 7, 7},
			expected: true,
		},
		{
			name:     "large input no duplicates",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: false,
		},
		{
			name:     "negative numbers with duplicates",
			input:    []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "mixed positive and negative no duplicates",
			input:    []int{-1, 0, 1, 2},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.input)

			if result != tt.expected {
				t.Errorf("containsDuplicate(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
