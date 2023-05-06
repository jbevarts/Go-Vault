package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subarraySum(t *testing.T) {
	assert.Equal(t, 2, subarraySum([]int{1, 1, 1}, 2))
	assert.Equal(t, 2, subarraySum([]int{1, 2, 3}, 2))
}
