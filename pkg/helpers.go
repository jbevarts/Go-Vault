package pkg

import "fmt"

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

	// if we are at start or last element is first element, return.
	if b.tail <= 0 {
		return
	}

	// new root becomes last element
	b.h = append([]T{b.h[b.tail]}, b.h...)
	fmt.Println("not heap")

	curIndex := 0
	for !b.isHeap(curIndex) {
		fmt.Println("not heap")
		curIndex = b.nextRoot(curIndex)
	}
}

func (b *binaryHeap[T]) isHeap(i int) bool {

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
