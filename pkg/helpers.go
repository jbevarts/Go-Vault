package pkg

func kahnsAlgorithm() ([]int, error) {

	return []int{}, nil
}

type topoSortVertex struct {
	seen     bool
	seeing   bool
	val      int
	outEdges []topoSortVertex
}

func topologicalSortDFS(e []int) ([]int, error) {

	// create vertices, map of val -> vertex.
	//

	// could keep vertices in a priority queue by whether or not seen is true.

	return []int{}, nil
}

// binaryHeap stores a collection of Node T's and maintains the Heap property.  Siblings have no orderings.
// cmp should be a comparator function that returns true when the first element is equal or higher priority
// to the second param.
type binaryHeap[T any] struct {
	h    []T
	tail int
	cmp  func(a, b T) bool
}

func NewBinaryHeap[T any](cmpFnc func(a, b T) bool) binaryHeap[T] {
	return binaryHeap[T]{
		h:    []T{},
		tail: -1, // tail is 0th position after insert
		cmp:  cmpFnc,
	}
}

// places last element at root and sinks into position to satisfy the heap property
func (b *binaryHeap[T]) heapify() {

	// new root becomes last element
	b.h[0] = b.h[b.tail]

	// last element decrements 1 position
	b.tail = b.tail - 1

	curIndex := 0
	for !b.isHeap(curIndex) {
		curIndex = b.nextRoot(curIndex)
	}
}

func (b *binaryHeap[T]) isHeap(i int) bool {

	// left child is 2n, right child is 2n + 1
	root := b.h[i]
	left := b.h[2*i]
	right := b.h[2*i+1]

	return b.cmp(root, left) && b.cmp(root, right)
}

func (b *binaryHeap[T]) nextRoot(i int) int {

	if b.cmp(b.h[2*i], b.h[2*i+1]) {
		return 2*i + 1
	} else {
		return 2 * i
	}
}

func (b *binaryHeap[T]) pop() T {

	// return first element
	t := b.h[0]

	// satisfy heap property after pop.
	b.heapify()

	return t
}

func (b *binaryHeap[T]) peek() T {

	return b.h[b.tail]
}

func (b *binaryHeap[T]) push(t T) {

	// add element to end of heap
	b.h = append(b.h, t)
	b.tail = b.tail + 1

	b.heapify()
}
