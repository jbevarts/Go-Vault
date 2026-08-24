package datastructures

import (
	"fmt"
	"maps"
	"slices"
)

// Backed by generic adjacency list
// If undirected, stores each edge twice, u->v & v->u
// If directed, u->v is a different edge than v->u if exists
type Graph[T comparable] struct {
	adj map[T][]T
}

// Directed acyclic graph checker
// implements topological stort
func IsGraphDAG[T comparable](g *Graph[T]) error {
	// for each 0-degree node
	// zeros, totalNodes := getZeroDegreeNodesInGraph(g)
	// if len(zeros) == 0 {
	// 	return errors.New("No zero degree edges")
	// }

	// validate dfs completes without hitting same node twice.
	var err error
	for k := range g.adj {
		err = preorderGraphCycleCheck(g, k, make(map[T]bool), func(val T, curPath map[T]bool) error {
			if curPath[val] {
				return fmt.Errorf("Not a DAG; cycle detected at %v", val)
			}
			curPath[val] = true
			fmt.Println(curPath)
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func preorderGraphCycleCheck[T comparable](g *Graph[T], root T, curPath map[T]bool, doFunc func(T, map[T]bool) error) error {
	if g == nil {
		return nil
	}

	if err := doFunc(root, curPath); err != nil {
		return err
	}
	fmt.Println(root, g.adj[root])

	for _, v := range g.adj[root] {
		if err := preorderGraphCycleCheck(g, v, maps.Clone(curPath), doFunc); err != nil {
			return err
		}
	}

	return nil
}

func getZeroDegreeNodesInGraph[T comparable](g *Graph[T]) ([]T, []T) {
	zeros := make(map[T]bool)

	for k := range g.adj {
		zeros[k] = true
	}
	for k := range g.adj {
		for _, v := range g.adj[k] {
			zeros[v] = false
		}
	}

	var result []T
	for val, isZero := range zeros {
		if isZero {
			result = append(result, val)
		}
	}

	return result, slices.Collect(maps.Keys(zeros))
}

func GetPreorderFromDirectedGraph[T comparable](g *Graph[T], root T) []T {
	if g == nil {
		return nil
	}

	var result []T
	visited := make(map[T]bool)
	getPreorderFromDirectedGraph(g, root, func(val T) {
		if visited[val] {
			return
		}
		result = append(result, val)
		visited[val] = true
	})

	return result
}

func getPreorderFromDirectedGraph[T comparable](g *Graph[T], root T, doFunc func(T)) {
	if g == nil {
		return
	}

	doFunc(root)
	for _, v := range g.adj[root] {
		getPreorderFromDirectedGraph(g, v, doFunc)
	}
}
