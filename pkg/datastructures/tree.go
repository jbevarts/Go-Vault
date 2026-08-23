package datastructures

import (
	"fmt"
	"strings"
)

/*
	Properties of a Tree:
	- connected
	- acyclic
	- undirected graph
	- unique path between every vertex
	- 1 root
	- N vertices, N-1 edges

	Types of trees:
	- Binary search tree
	- Binary tree
	- AVL tree
	- B-tree
	- Red–black tree
	- Splay tree
	- Self-balancing binary search tree
	- M-ary tree

*/

type TreeNode[T comparable] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTreeNode[T comparable](val T, left *TreeNode[T], right *TreeNode[T]) *TreeNode[T] {
	return &TreeNode[T]{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func GetPreorderIntegersFromBST(t *TreeNode[int]) []int {
	if t == nil {
		return nil
	}

	var result []int
	getPreorderIntegersFromBST(t, func(v int) {
		result = append(result, v)
	})

	return result
}

func getPreorderIntegersFromBST(t *TreeNode[int], doFunc func(int)) {
	if t == nil {
		return
	}

	doFunc(t.Val)
	getPreorderIntegersFromBST(t.Left, doFunc)
	getPreorderIntegersFromBST(t.Right, doFunc)
}

func GetInorderIntegersFromBST(t *TreeNode[int]) []int {
	// heapindex layout, 2n+1/2n+2
	if t == nil {
		return nil
	}

	var list []int
	getInorderIntegersFromBST(t, func(v int) {
		list = append(list, v)
	})

	return list
}

func getInorderIntegersFromBST(t *TreeNode[int], doFunc func(int)) {
	if t == nil {
		return
	}

	getInorderIntegersFromBST(t.Left, doFunc)
	doFunc(t.Val)
	getInorderIntegersFromBST(t.Right, doFunc)

}

func GenerateTreeFromIntegers(vals []int) *TreeNode[int] {
	// vals should be a valid string representing a tree where for each position i in the array
	// i == parent node, 2i+1 is left child, 2i+2 is right child.
	// Example: 1,2,3,4,5
	// for 1, left is 2(0)+1 == "2", right is 2(0)+2 == "3"
	// for 2, left is 2(1)+1 == "4", right is 2(1)+2 == "5"
	// so, Tree looks like:
	//   1
	//  2 3
	// 4 5
	if len(vals) == 0 {
		return nil
	}
	return generateTreeFromIntegers(vals, 0)
}

func generateTreeFromIntegers[T comparable](vals []T, index int) *TreeNode[T] {
	t := NewTreeNode(vals[index], nil, nil)

	leftIndex := 2*index + 1
	rightIndex := 2*index + 2

	if len(vals) > leftIndex {
		t.Left = generateTreeFromIntegers(vals, leftIndex)
	}

	if len(vals) > rightIndex {
		t.Right = generateTreeFromIntegers(vals, rightIndex)
	}
	return t
}

// PrintTree writes a directory-style visualization of t to stdout.
// Left child is printed first; a missing left child is shown as "." when a right child exists.
func PrintTree[T comparable](t *TreeNode[T]) {
	fmt.Print(formatTree(t))
}

func (t *TreeNode[T]) String() string {
	return formatTree(t)
}

func formatTree[T comparable](t *TreeNode[T]) string {
	if t == nil {
		return "<nil>\n"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%v\n", t.Val)
	writeChildren(t, "", &b)
	return b.String()
}

func writeChildren[T comparable](t *TreeNode[T], prefix string, b *strings.Builder) {
	if t.Left == nil && t.Right == nil {
		return
	}

	writeChild := func(child *TreeNode[T], last bool) {
		branch, next := "├── ", prefix+"│   "
		if last {
			branch, next = "└── ", prefix+"    "
		}
		if child == nil {
			b.WriteString(prefix)
			b.WriteString(branch)
			b.WriteString(".\n")
			return
		}
		b.WriteString(prefix)
		b.WriteString(branch)
		fmt.Fprintf(b, "%v\n", child.Val)
		writeChildren(child, next, b)
	}

	if t.Right == nil {
		writeChild(t.Left, true)
		return
	}
	writeChild(t.Left, false)
	writeChild(t.Right, true)
}
