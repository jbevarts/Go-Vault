package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binaryHeap(t *testing.T) {

	bh := NewBinaryHeap[int](func(a, b int) bool {
		return a <= b
	})

	bh.push(1)
	bh.push(15)
	bh.push(16)
	bh.push(4)
	bh.push(2)

	assert.Equal(t, bh.pop(), 1)
	assert.Equal(t, bh.pop(), 2)
	assert.Equal(t, bh.pop(), 4)
	assert.Equal(t, bh.pop(), 15)
	assert.Equal(t, bh.pop(), 16)

	bh = NewBinaryHeap[int](func(a, b int) bool {
		return a >= b
	})

	bh.push(1)
	bh.push(15)
	bh.push(16)
	bh.push(4)
	bh.push(2)

	assert.Equal(t, bh.pop(), 16)
	assert.Equal(t, bh.pop(), 15)
	assert.Equal(t, bh.pop(), 4)
	assert.Equal(t, bh.pop(), 2)
	assert.Equal(t, bh.pop(), 1)

	bh = NewBinaryHeap[int](func(a, b int) bool {
		return a >= b
	})

	bh.push(1)
	bh.push(15)
	assert.Equal(t, bh.pop(), 15)
	bh.push(2)
	bh.push(6)
	assert.Equal(t, bh.pop(), 6)
	bh.push(3)
	assert.Equal(t, bh.pop(), 3)
	assert.Equal(t, bh.pop(), 2)
	assert.Equal(t, bh.pop(), 1)

}
