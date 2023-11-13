package leetcode

import (
	"fmt"
)

// This solution could integrate the usage of a Set but it would take up more memory.

type void struct{}
var member void

func isValidSudoku(board [][]byte) bool {
    // build columns
    // for each row, validate
    // for each column, validate
    // build rows out of each subsquare
    // for those 9 additional rows, validate.
    // total row validations: 27

    total := append(board, makeColumns(board)...)
    total = append(total, makeSquares(board)...)

    for _, v := range total {
        v := isValid(v)
        if !v {
            return false
        }
    }
    return true
}

func makeColumns(elems [][]byte) [][]byte{
    res := make([][]byte, 9)
    for i := range elems { // constant 9
        for _, v := range elems {
            res[i] = append(res[i], v[i])
        }
    }
    return res
}
// scan each row, every 3 rows is 3 squares
// every 3 elements, change square
// continue until 3 sqares built and append square to final square list
func makeSquares(elems [][]byte) [][]byte{
    res := make([][]byte, 9) 
    boxMin := 0
    for k, v := range elems {
        if k != 0 && k%3 == 0 {
            boxMin = k
        }
        box := boxMin
        for i, j := range v {
            if i != 0 && i % 3 == 0 {
                box = box + 1
            }
            if string(j) != "." {
                res[box] = append(res[box], j)
            }
        }
    }
    
    return res
}

// Assumptions: len() == 9, elem == 1-9 or "."
func isValid(elems []byte) bool {
    m := make(map[string]void, 0)
    for _, v := range elems {
        val := string(v)
        if val == "." {
            continue
        }
        if _, ok := m[val]; ok {
            printBytes(elems)
            return false
        }
        m[val] = member
    }
    return true
}

func printBytes(bts []byte) {
    var b []string
    for _, v := range bts {
        b = append(b, string(v))
    } 
    fmt.Println(b)
}