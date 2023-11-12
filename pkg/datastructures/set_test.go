package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Set(t *testing.T) {

	set := NewSet[int]()
	set.Insert(2)
	assert.Equal(t, true, set.Contains(2))
	assert.Equal(t, []int{2}, set.ListElements())
	set.Insert(2)
	assert.Equal(t, true, set.Contains(2))
	assert.Equal(t, []int{2}, set.ListElements())
	set.Remove(2)
	assert.Equal(t, false, set.Contains(2))
	assert.Equal(t, []int{}, set.ListElements())
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)
	set.Insert(3)
	set.Insert(4)
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, true, set.Contains(2))
	assert.Equal(t, true, set.Contains(3))
	assert.Equal(t, true, set.Contains(4))
	assert.Equal(t, false, set.Contains(5))
	set.Remove(2)
	set.Remove(5)
	assert.Equal(t, true, set.Contains(1))
	assert.Equal(t, false, set.Contains(2))
	assert.Equal(t, true, set.Contains(3))
	assert.Equal(t, true, set.Contains(4))
	assert.Equal(t, false, set.Contains(5))
	assert.Equal(t, 3, len(set.ListElements()))
}
