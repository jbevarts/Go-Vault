package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A subsequence is an array that can be derived from another array by deleting some
// or no elements without changing the order of the remaining elements.

func Test_lengthOfLIS(t *testing.T) {
	assert.Equal(t, 1, lengthOfLIS([]int{1}))
	assert.Equal(t, 2, lengthOfLIS([]int{1, 2}))
	assert.Equal(t, 3, lengthOfLIS([]int{1, 2, 0, 3}))
	assert.Equal(t, 3, lengthOfLIS([]int{1, 2, 3, 0}))
	assert.Equal(t, 3, lengthOfLIS([]int{1, 0, 2, 3}))
	assert.Equal(t, 4, lengthOfLIS([]int{0, 1, 2, 3}))
	assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Equal(t, 4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	assert.Equal(t, 1, lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
