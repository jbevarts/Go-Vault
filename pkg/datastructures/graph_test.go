package datastructures

import (
	"errors"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getZeroDegreeNodesInGraph(t *testing.T) {
	tt := []struct {
		name       string
		graph      *Graph[int]
		totalNodes int
		expected   []int
	}{
		{
			name: "test",
			graph: &Graph[int]{
				adj: map[int][]int{
					1: {2, 3, 4},
					2: {3, 5},
					5: {8, 9},
				},
			},
			totalNodes: 7,
			expected:   []int{1},
		},
		{
			name: "test2",
			graph: &Graph[int]{
				adj: map[int][]int{
					1: {2, 3, 4},
					2: {3},
					5: {8, 9},
				},
			},
			totalNodes: 7,
			expected:   []int{1, 5},
		},
	}

	for _, tc := range tt {
		got, totalNodes := getZeroDegreeNodesInGraph(tc.graph)
		slices.Sort(got)
		assert.Equal(t, tc.expected, got)
		assert.Equal(t, tc.totalNodes, totalNodes)
	}
}

func Test_IsGraphDAG(t *testing.T) {
	tt := []struct {
		name          string
		graph         *Graph[int]
		expectedError error
	}{
		{
			name: "test",
			graph: &Graph[int]{
				adj: map[int][]int{
					1: {2, 3, 4},
					2: {3, 5},
					5: {8, 9},
				},
			},
			expectedError: nil,
		},
		{
			name: "test2",
			graph: &Graph[int]{
				adj: map[int][]int{
					1: {2, 3, 4},
					2: {4, 5},
					4: {1},
					5: {8, 9},
				},
			},
			expectedError: errors.New("Not a DAG; cycle detected at 4"),
		},
		{
			name: "test3",
			graph: &Graph[int]{
				adj: map[int][]int{
					1: {2, 3, 4},
					2: {4, 5},
					4: {2},
					5: {8, 9},
				},
			},
			expectedError: errors.New("Not a DAG; cycle detected at 2"),
		},
		{
			name: "test4",
			graph: &Graph[int]{
				adj: map[int][]int{
					1: {2},
					2: {1},
					3: {4},
				},
			},
			expectedError: errors.New("Not a DAG; cycle detected at 2"),
		},
	}

	for _, tc := range tt {
		got := IsGraphDAG(tc.graph)
		assert.Equal(t, tc.expectedError, got, tc.name)
	}
}

func Test_GetPreorderFromDirectedGraph(t *testing.T) {
	tt := []struct {
		name          string
		directedGraph *Graph[int]
		root          int
		expected      []int
	}{
		{
			name: "test",
			directedGraph: &Graph[int]{
				adj: map[int][]int{
					1: {2, 3, 4},
					2: {3, 5, 6},
					5: {8, 10},
				},
			},
			root:     2,
			expected: []int{2, 3, 5, 8, 10, 6},
		},
		{
			name: "test2",
			directedGraph: &Graph[int]{
				adj: map[int][]int{
					1: {2, 3, 4},
					2: {3, 5, 6},
					5: {8, 10},
				},
			},
			root:     1,
			expected: []int{1, 2, 3, 5, 8, 10, 6, 4},
		},
	}

	for _, tc := range tt {
		got := GetPreorderFromDirectedGraph(tc.directedGraph, tc.root)
		assert.Equal(t, tc.expected, got, "%v\n does not equal\n v", got, tc.expected)
	}
}
