package datastructures

import (
	"testing"
)

func Test_GenerateTreeFromIntegers(t *testing.T) {
	tc := []struct {
		name     string
		integers []int
		expected *TreeNode[int]
	}{
		{
			name:     "test",
			integers: []int{1, 2, 3, 4, 5, 6, 7},
			expected: nil,
		},
	}

	for _, tt := range tc {
		PrintTree(GenerateTreeFromIntegers(tt.integers))
	}
}
