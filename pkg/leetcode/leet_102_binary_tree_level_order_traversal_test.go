package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {

	n := &TreeNode{Val: 3}
	n.Left = &TreeNode{Val: 9}
	n.Right = &TreeNode{Val: 20}
	n.Right.Right = &TreeNode{Val: 7}
	n.Right.Left = &TreeNode{Val: 15}

	response := levelOrder(n)

	if response == nil {
		t.Error("nil response not expected")
	}
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	for i, v := range response {
		for _, w := range expected[i] {
			assert.Equal(t, w, v)
		}
	}
}
