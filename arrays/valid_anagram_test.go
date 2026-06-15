package arrays

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "valid anagram",
			s:        "listen",
			t:        "silent",
			expected: true,
		},
		{
			name:     "different lengths",
			s:        "abc",
			t:        "ab",
			expected: false,
		},
		{
			name:     "not an anagram",
			s:        "hello",
			t:        "world",
			expected: false,
		},
		{
			name:     "same string",
			s:        "test",
			t:        "test",
			expected: true,
		},
		{
			name:     "empty strings",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "single character same",
			s:        "a",
			t:        "a",
			expected: true,
		},
		{
			name:     "single character different",
			s:        "a",
			t:        "b",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)

			if result != tt.expected {
				t.Errorf("isAnagram(%q, %q) = %v; expected %v", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
