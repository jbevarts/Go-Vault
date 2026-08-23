package algorithms

import (
	ds "Go-Blind/pkg/datastructures"
)

// Uses TreeNode to return array of integers of inorder traversal of a tree.

func inorderString(t *ds.TreeNode[int]) []int {
	var result []int
	inOrder(t, func() {
		result = append(result, t.Val)
	})
	return result
}

func inOrder(t *ds.TreeNode[int], doFunc func()) {
	if t == nil {
		return
	}

	inOrder(t.Left, doFunc)
	doFunc()
	inOrder(t.Right, doFunc)
}
