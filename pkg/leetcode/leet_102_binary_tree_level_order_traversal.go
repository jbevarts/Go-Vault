package pkg

// Given the root of a binary tree, ( which is not strictly )
// return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// Idea is to store each row's values in a FIFO queue and then iterate over the queue,
	// adding each elements children to the queue.  Each time we pop a value from the queue, we visit it
	// and append the current levels slice.

	// need to implement FIFO queue.
	// TreeNode would be useful as well.

	return nil

}
