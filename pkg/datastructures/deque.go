package datastructures

// Deque is a double-ended queue that can be inserted and removed from both ends.

type Deque[T any] struct {
	deque        []T
	head         int
	tail         int
	loadCapacity int
}

type Dequer[T any] interface {
	PushHead()
	PopHead() T
	PushTail()
	PopTail() T
	Size() int
}

func (d *Deque[T]) PushHead(ele T) {}
func (d *Deque[T]) PopHead() *T {
	return nil
}
func (d *Deque[T]) PushTail(ele T) {}

// Return value is nil when non-existent.
func (d *Deque[T]) PopTail() *T {

	if d.Size() < 1 {
		return nil
	}

	absTail := d.tail
	if d.tail < 0 {
		absTail = len(d.deque) + d.tail
	}

	res := d.deque[d.tail]
	// Only occurs when dequeu has a single valid element.
	if absTail == d.head {
		return &res
	}

	d.tail--
	return &res

}
func (d *Deque[T]) Size() int {

	absTail := d.tail
	if d.tail < 0 {
		absTail = len(d.deque) + d.tail
	}

	return absTail - d.head
}
