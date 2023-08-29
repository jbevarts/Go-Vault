package pkg

import (
	"Go-Blind/pkg/datastructures"
)

// Given the root of a binary tree, ( which is not strictly )
// return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
// One caveat to this problem is that you must preserve the levels in your put.  If a level has 8 children
// that entry will be length 8.  This challenges the traditional BFS approach because keeping track of an end of a row
// is difficult without adding nil's into your queue which grows in size annoyingly.

// I believe the best answer is actually recursive BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	return helper(datastructures.NewDeque[TreeNode](10), [][]int{})

}

func helper(d datastructures.Deque[TreeNode], final [][]int) [][]int {

	if d.Size() == 0 {
		return final
	}

	var nextDeque datastructures.Deque[TreeNode]
	cur := []int{}

	for d.Size() > 0 {
		next, _ := d.PopHead()
		cur = append(cur, next.Val)
		pushNodeToTail(&nextDeque, next.Left)
		pushNodeToTail(&nextDeque, next.Right)
	}

	final = append(final, cur)

	return helper(nextDeque, final)
}

func pushNodeToTail(d *datastructures.Deque[TreeNode], n *TreeNode) {
	if n != nil {
		d.PushTail(*n)
	}
}
