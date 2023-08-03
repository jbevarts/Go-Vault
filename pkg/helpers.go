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

	// create vertices

	// could keep vertices in a priority queue by whether or not seen is true.

	return []int{}, nil
}

// i need a priority queue based on an arbitrary priority.
// perhaps it takes a function and a generic and is able to run the function(s) on the generic.
