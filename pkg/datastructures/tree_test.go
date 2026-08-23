package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateBSTFromIntegers(t *testing.T) {
	tc := []struct {
		name     string
		integers []int
		expected *TreeNode[int]
	}{
		{
			name:     "test",
			integers: []int{4, 2, 6, 1, 3, 5, 7}, // assumes heap-index ordering, 2n+1/2n+1
			expected: &TreeNode[int]{
				Val: 4,
				Left: &TreeNode[int]{
					Val: 2,
					Left: &TreeNode[int]{
						Val: 1,
					},
					Right: &TreeNode[int]{
						Val: 3,
					},
				},
				Right: &TreeNode[int]{
					Val: 6,
					Left: &TreeNode[int]{
						Val: 5,
					},
					Right: &TreeNode[int]{
						Val: 7,
					},
				},
			},
		},
	}

	for _, tt := range tc {
		got := GenerateTreeFromIntegers(tt.integers)
		assert.Equal(t, tt.expected, got, "%v does not equal %v", tt.expected, got)
	}
}

func Test_GetInorderIntegersFromBST(t *testing.T) {
	tc := []struct {
		name     string
		tree     *TreeNode[int] // must be valid BST
		expected []int
	}{
		{
			name: "test",
			tree: &TreeNode[int]{
				Val:   2,
				Left:  &TreeNode[int]{Val: 1},
				Right: &TreeNode[int]{Val: 3},
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tc {
		got := GetInorderIntegersFromBST(tt.tree)
		assert.Equal(t, tt.expected, got, fmt.Sprintf("%v doesn't equal %v", got, tt.expected))
	}
}

func Test_GetPreorderIntegersFromBST(t *testing.T) {
	tc := []struct {
		name     string
		tree     *TreeNode[int] // must be valid BST
		expected []int
	}{
		{
			name: "test",
			tree: &TreeNode[int]{
				Val:   2,
				Left:  &TreeNode[int]{Val: 1},
				Right: &TreeNode[int]{Val: 3},
			},
			expected: []int{2, 1, 3},
		},
	}

	for _, tt := range tc {
		got := GetPreorderIntegersFromBST(tt.tree)
		assert.Equal(t, tt.expected, got, fmt.Sprintf("%v doesn't equal %v", got, tt.expected))
	}
}
