package datastructures

// generic Set implementation with a map backing.

type void struct{}
var member void

type Set[T comparable] struct {
	elems map[T]void
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elems: make(map[T]void, 0)}
}

func (s *Set[T]) Insert(elem T) {
	s.elems[elem] = member
} 

func (s *Set[T]) Remove(elem T) {
	delete(s.elems, elem)
}

func (s *Set[T]) Contains(elem T) bool {
	if _, ok := s.elems[elem]; ok {
		return true
	}
	return false
}

// Non-deterministic output
func (s *Set[T]) ListElements() []T {
	elems := []T{}
	for k, _ := range s.elems {
		elems = append(elems, k)
	}
	return elems
}