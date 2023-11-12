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

	// If we have completed the entire of deque of nodes, return the level ordering, final.
	if d.Size() == 0 {
		return final
	}

	// Create a new deque for each row.
	var nextDeque datastructures.Deque[TreeNode]
	// Maintain 'this rows' ordering with cur.
	cur := []int{}

	// d should have size == qty of previous row's children.
	for d.Size() > 0 {
		// Take next child
		next, _ := d.PopHead()
		
		// Add child to the end of this rows ordering.
		cur = append(cur, next.Val)

		// for both children of next, add to end of queue to be processed by next call. 
		pushNodeToTail(&nextDeque, next.Left)
		pushNodeToTail(&nextDeque, next.Right)
	}

	// Concatenate cur with previous calls level order.
	final = append(final, cur)

	// Recursively call helper with deque of children and partial answer, final.
	return helper(nextDeque, final)
}

func pushNodeToTail(d *datastructures.Deque[TreeNode], n *TreeNode) {

	// Avoid populating nil children.
	if n != nil {
		d.PushTail(*n)
	}
}
