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

func (b *binaryHeap[T]) sink() {

	// if we are at start or last element is first element, return.
	if b.tail <= 0 {
		return
	}

	curIndex := 0
	for !b.isSinkHeap(curIndex) {
		newHead := b.nextRoot(curIndex)

		tmp := b.h[curIndex]
		b.h[curIndex] = b.h[newHead]
		b.h[newHead] = tmp

		curIndex = newHead
	}
}

func (b *binaryHeap[T]) bubbleUp() {

	// if we are at start or last element is first element, return.
	if b.tail <= 0 {
		return
	}

	curIndex := b.tail
	// only need to compare new node with parent, repeatedly.
	for !b.isBubbleHeap(curIndex) {
		var newHead int

		if curIndex%2 == 0 {
			newHead = curIndex/2 - 1
		} else {
			// assuming round down
			newHead = curIndex / 2
		}

		tmp := b.h[newHead]
		b.h[newHead] = b.h[curIndex]
		b.h[curIndex] = tmp

		curIndex = newHead
	}
}

func (b *binaryHeap[T]) isSinkHeap(i int) bool {

	// left child is 2n + 1, right child is 2n + 2 ( when zero based )
	root := b.h[i]

	if b.tail >= 2*i+1 {
		left := b.h[2*i+1]

		if b.tail >= 2*i+2 {
			right := b.h[2*i+2]
			return b.cmp(root, left) && b.cmp(root, right)
		}
		return b.cmp(root, left)
	}

	// leaf node is always a heap
	return true
}

func (b *binaryHeap[T]) isBubbleHeap(i int) bool {

	if i == 0 {
		return true
	}

	node := b.h[i]

	if i%2 == 0 {
		return b.cmp(b.h[i/2-1], node)
	} else {
		// assuming round down
		return b.cmp(b.h[i/2], node)
	}
}

func (b *binaryHeap[T]) nextRoot(i int) int {

	if b.tail >= 2*i+1 {
		left := b.h[2*i+1]

		if b.tail >= 2*i+2 {
			right := b.h[2*i+2]

			if b.cmp(left, right) {
				return 2*i + 1
			} else {
				return 2*i + 2
			}
		}
		return 2*i + 1
	}

	return b.tail
}

func (b *binaryHeap[T]) pop() T {

	// return first element
	t := b.h[0]

	// satisfy heap property after pop.
	b.h[0] = b.h[b.tail]
	b.h = b.h[:len(b.h)-1]

	b.tail = b.tail - 1

	b.sink()

	return t
}

func (b *binaryHeap[T]) peek() T {

	return b.h[0]
}

func (b *binaryHeap[T]) push(t T) {

	// add element to bottom of heap
	b.h = append(b.h, t)
	b.tail = b.tail + 1

	b.bubbleUp()
}
