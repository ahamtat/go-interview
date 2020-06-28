package counter

import (
	"testing"
)

func TestCountGoWords(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		expected int
	}{
		{
			name:     "Single Go",
			text:     " Go   ",
			expected: 1,
		},
		{
			name:     "Three Go",
			text:     "Go Go Go",
			expected: 3,
		},
		{
			name:     "Google",
			text:     "Google",
			expected: 0,
		},
		{
			name:     "Golang Gopher",
			text:     "Golang Gopher",
			expected: 0,
		},
		{
			name:     "Quoted",
			text:     `Word "Go" is awesome`,
			expected: 1,
		},
		{
			name:     "Uppercase",
			text:     `"GO"`,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n := CountGoWords(tc.text)
			if n != tc.expected {
				t.Errorf("wrong words number = %d, expected %d words", n, tc.expected)
			}
		})
	}
}
