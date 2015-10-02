package main

import (
	"fmt"
)

// add takes in two matrices and returns their sum
func add(mat1, mat2 Matrix, c chan Matrix) {
	dim1, dim2 := len(mat1.Val[0]), len(mat2.Val[0])
	if dim1 != dim2 {
		c <- Matrix{e: fmt.Errorf("Cannot add matrices of varying lengths: %v != %v", dim1, dim2)}
	} else if dim1 == 0 {
		c <- Matrix{e: fmt.Errorf("Empty matrices.")}
	}

	sum := make([][]Element, dim1)
	for i := 0; i < dim1; i++ {
		sum[i] = make([]Element, dim1)
		for j := 0; j < dim1; j++ {
			sum[i][j] = Element{mat1.Val[i][j].Val + mat2.Val[i][j].Val}
		}
	}
	c <- Matrix{Val: sum}
}
