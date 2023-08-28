package datastructures

// Deque is a self-resizing, double-ended queue that can be inserted and removed from both ends.

type Deque[T any] struct {
	deque        []T
	head         int
	tail         int
	loadCapacity int
	size int
}

func NewDeque[T any](capacity int) *Deque[T] {
	return &Deque[T]{
		deque: make([]T, capacity),
	}
}

func (d *Deque[T]) resizeDeque() {

	n := make([]T, d.loadCapacity * 2)
	tempHead := d.head

	for i := range d.deque {
		if d.normalizePosition(tempHead) == d.normalizePosition(d.tail) {
			return
		}
		n[i] = d.deque[d.normalizePosition(tempHead)]
		tempHead++
	}
	d.head = 0
	d.tail = d.size - 1
}

func (d *Deque[T]) PushHead(ele T) {

	if d.normalizePosition(d.head - 1) == d.normalizePosition(d.tail) {
		d.resizeDeque()
	}

	d.head--
	d.size++
	d.deque[d.normalizePosition(d.head)] = ele
}
func (d *Deque[T]) PopHead() *T {

	if d.size < 1 {
		return nil
	}

	d.head++
	d.size--
	return &d.deque[d.normalizePosition(d.head - 1)]
}
func (d *Deque[T]) PushTail(ele T) {

	if d.normalizePosition(d.tail + 1) == d.normalizePosition(d.head) {
		d.resizeDeque()
	}

	d.tail++
	d.size++
	d.deque[d.normalizePosition(d.tail)] = ele
}

func (d *Deque[T]) normalizePosition(i int) int {
	// -1 == last element
	// -2 == second to last element

	// Len() == 1st element
	// Len() + 1 == 2nd element

	if i < 0 {
		return len(d.deque) + i 
	}

	if i >= len(d.deque) {
		return i - len(d.deque)
	}

	return i
}

// Return value is nil when non-existent.
func (d *Deque[T]) PopTail() *T {

	if d.size < 1 {
		return nil
	}

	d.tail--
	d.size--
	return &d.deque[d.tail + 1]

}

func (d *Deque[T]) PeekTail() *T {
	return &d.deque[d.tail]
}

func (d *Deque[T]) PeekHead() *T {
	return &d.deque[d.head]
}