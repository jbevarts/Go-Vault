package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func tryOrFail(t *testing.T, try func() (int, error)) int {

	r, err := try()
	if err != nil {
		t.Fail()
	}
	return r
}
func Test_deque(t *testing.T) {

	bh := NewDeque[int](4)

	bh.PushTail(1)
	bh.PushTail(15)
	bh.PushTail(16)
	bh.PushTail(4)
	fmt.Println(bh.deque)

	bh.PushTail(2)

	fmt.Println(bh.deque)

	assert.Equal(t, tryOrFail(t, bh.PopTail), 2)
	assert.Equal(t, tryOrFail(t, bh.PopTail), 4)
	assert.Equal(t, tryOrFail(t, bh.PopTail), 16)
	assert.Equal(t, tryOrFail(t, bh.PopTail), 15)
	assert.Equal(t, tryOrFail(t, bh.PopTail), 1)
	fmt.Println(bh.deque)

	bh.PushTail(1)
	bh.PushTail(15)
	bh.PushTail(16)
	bh.PushTail(4)
	bh.PushTail(2)
	fmt.Println(bh.deque)

	assert.Equal(t, tryOrFail(t, bh.PopHead), 1)
	assert.Equal(t, tryOrFail(t, bh.PopHead), 15)
	assert.Equal(t, tryOrFail(t, bh.PopHead), 16)
	assert.Equal(t, tryOrFail(t, bh.PopHead), 4)
	assert.Equal(t, tryOrFail(t, bh.PopHead), 2)

	fmt.Println(bh.deque)

	bh.PushHead(1)
	bh.PushHead(15)
	bh.PushHead(16)
	bh.PushHead(4)
	bh.PushHead(2)

	fmt.Println(bh.deque)

	assert.Equal(t, tryOrFail(t, bh.PopHead), 2)
	assert.Equal(t, tryOrFail(t, bh.PopHead), 4)
	assert.Equal(t, tryOrFail(t, bh.PopHead), 16)
	assert.Equal(t, tryOrFail(t, bh.PopHead), 15)
	assert.Equal(t, tryOrFail(t, bh.PopHead), 1)

	fmt.Println(bh.deque)

	bh.PushHead(1)
	bh.PushHead(15)
	bh.PushHead(16)
	bh.PushHead(4)
	bh.PushHead(2)

	fmt.Println(bh.deque)

	assert.Equal(t, tryOrFail(t, bh.PopTail), 1)
	assert.Equal(t, tryOrFail(t, bh.PopTail), 15)
	assert.Equal(t, tryOrFail(t, bh.PopTail), 16)
	assert.Equal(t, tryOrFail(t, bh.PopTail), 4)
	assert.Equal(t, tryOrFail(t, bh.PopTail), 2)

	fmt.Println(bh.deque)

}
