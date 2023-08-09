package algorithms

type topoSortVertex struct {
	seen     bool
	seeing   bool
	val      int
}

// func topoSortDFS(adj [][]int) ([]int, error) {

// 	if len(adj) < 1 {
// 		return []int, nil
// 	}

// 	var arr []topoSortVertex
// 	for k, v := range adj {
// 		arr = append(arr, topoSortVertex{val: v})
// 	}

// 	for k, v := range adj {

// 	}

// 	// create vertices, map of val -> vertex.
// 	//

// 	// could keep vertices in a priority queue by whether or not seen is true.

// 	return []int{}, nil
// }
