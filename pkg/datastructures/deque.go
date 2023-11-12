package datastructures

import (
	"errors"
)

// Deque is a self-resizing, double-ended queue that can be inserted and removed from both ends.

type Deque[T any] struct {
	deque        []T
	head         int
	tail         int
	loadCapacity int
	size         int
}

func NewDeque[T any](capacity int) Deque[T] {
	return Deque[T]{
		deque:        make([]T, capacity),
		loadCapacity: capacity,
	}
}

func NewDequeWithValues[T any](vals []T, capacity int) Deque[T] {

	return Deque[T]{
		deque:        vals,
		loadCapacity: capacity,
	}
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) resizeDeque() {

	n := make([]T, d.loadCapacity*2)
	tempHead := d.head

	for i := 0; i < cap(d.deque); i++ {
		n[i] = d.deque[d.normalizePosition(tempHead)]
		tempHead++
	}
	d.head = 0
	d.tail = d.size - 1
	d.deque = n
}

func (d *Deque[T]) PushHead(ele T) {

	if d.normalizePosition(d.head-1) == d.normalizePosition(d.tail) {
		d.resizeDeque()
	}
	d.size++

	if d.size == 1 {
		d.deque[d.normalizePosition(d.head)] = ele
		return
	}

	d.head--
	d.size++
	d.deque[d.normalizePosition(d.head)] = ele
}

func (d *Deque[T]) PopHead() (T, error) {

	if d.size < 1 {
		return d.deque[0], errors.New("nothing to pop")
	}
	d.size--

	if d.size == 0 {
		return d.deque[d.normalizePosition(d.head)], nil
	}

	d.head++
	return d.deque[d.normalizePosition(d.head-1)], nil

}
func (d *Deque[T]) PushTail(ele T) {

	if d.normalizePosition(d.tail+1) == d.normalizePosition(d.head) {
		d.resizeDeque()
	}
	d.size++

	if d.size > 1 {
		d.tail++
	}

	d.deque[d.normalizePosition(d.tail)] = ele
}

func (d *Deque[T]) normalizePosition(i int) int {
	// -1 == last element
	// -2 == second to last element

	// Len() == 1st element
	// Len() + 1 == 2nd element
	if i < 0 {
		return cap(d.deque) + i
	}

	if i >= cap(d.deque) {
		return i - cap(d.deque)
	}

	return i
}

// Return value is nil when non-existent.
func (d *Deque[T]) PopTail() (T, error) {

	if d.size < 1 {
		return d.deque[0], errors.New("nothing to pop")
	}
	d.size--

	if d.size == 0 {
		return d.deque[d.tail], nil
	}

	d.tail--
	return d.deque[d.normalizePosition(d.tail+1)], nil
}

func (d *Deque[T]) PeekTail() T {
	return d.deque[d.tail]
}

func (d *Deque[T]) PeekHead() T {
	return d.deque[d.head]
}
