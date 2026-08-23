package algorithms

import (
	"fmt"
	"testing"

	ds "Go-Blind/pkg/datastructures"
)

func Test_inorder(t *testing.T) {
	tc := []struct {
		name string
		tree *ds.TreeNode[int]
	}{
		{
			name: "test",
			tree: ds.NewTreeNode(1, nil, nil),
		},
	}

	for _, tt := range tc {
		fmt.Println(tt.name)
	}
}
