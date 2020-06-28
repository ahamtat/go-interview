package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BinarySearch(t *testing.T) {
	testCases := []struct {
		name          string
		array         []int
		item          int
		expectedIndex int
	}{
		{
			name:          "Test 1",
			array:         []int{1, 2, 3, 4, 6, 8, 10},
			item:          2,
			expectedIndex: 1,
		},
		{
			name:          "Test 2",
			array:         []int{1, 2, 3, 4, 6, 8, 10},
			item:          4,
			expectedIndex: 3,
		},
		{
			name:          "Test 3",
			array:         []int{1, 2, 3, 4, 6, 8, 10},
			item:          10,
			expectedIndex: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			index := BinarySearch(tc.array, 0, len(tc.array), tc.item)
			assert.Equal(t, tc.expectedIndex, index)
		})
	}
}
