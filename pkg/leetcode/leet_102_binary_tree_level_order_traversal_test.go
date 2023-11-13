package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {

	tests := []struct {
		tree *TreeNode
		want [][]int
	}{
		{
			tree: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20, 
						Left: &TreeNode{
							Val: 15,
						}, 
						Right: &TreeNode{
							Val: 7,
						},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			tree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 4,
					}, 
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 20, 
					Left: &TreeNode{
						Val: 17,
						Left: &TreeNode{
							Val: 1,
						}, 
						Right: &TreeNode{
							Val: 7,
						},
					}, 
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			want: [][]int{{3}, {9, 20}, {4, 2, 17, 8}},
		},
		{
			tree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
			},
			want: [][]int{{3}, {9}},
		},
		{
			tree: &TreeNode{
				Val: 3,
			},
			want: [][]int{{3}},
		},
		{
			tree: nil,
			want: [][]int{},
		},
		{
			tree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			want: [][]int{{3}, {9, 6}},
		},
	}

	for _, tt := range tests {
		res := levelOrder(tt.tree)

		if res == nil {
			t.Error("nil response not expected")
		}
		fmt.Println(tt.want)
		fmt.Println(res)
		for i, v := range res {
			for _, w := range tt.want[i] {
				assert.Equal(t, w, v)
			}
		}
	}
}
