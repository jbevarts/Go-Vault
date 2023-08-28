package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binaryHeap(t *testing.T) {

	bh := NewBinaryHeap[int](func(a, b int) bool {
		return a <= b
	})

	bh.Push(1)
	bh.Push(15)
	bh.Push(16)
	bh.Push(4)
	bh.Push(2)

	assert.Equal(t, bh.Pop(), 1)
	assert.Equal(t, bh.Pop(), 2)
	assert.Equal(t, bh.Pop(), 4)
	assert.Equal(t, bh.Pop(), 15)
	assert.Equal(t, bh.Pop(), 16)

	bh = NewBinaryHeap[int](func(a, b int) bool {
		return a >= b
	})

	bh.Push(1)
	bh.Push(15)
	bh.Push(16)
	bh.Push(4)
	bh.Push(2)

	assert.Equal(t, bh.Pop(), 16)
	assert.Equal(t, bh.Pop(), 15)
	assert.Equal(t, bh.Pop(), 4)
	assert.Equal(t, bh.Pop(), 2)
	assert.Equal(t, bh.Pop(), 1)

	bh = NewBinaryHeap[int](func(a, b int) bool {
		return a >= b
	})

	bh.Push(1)
	bh.Push(15)
	assert.Equal(t, bh.Pop(), 15)
	bh.Push(2)
	bh.Push(6)
	assert.Equal(t, bh.Pop(), 6)
	bh.Push(3)
	assert.Equal(t, bh.Pop(), 3)
	assert.Equal(t, bh.Pop(), 2)
	assert.Equal(t, bh.Pop(), 1)

}
