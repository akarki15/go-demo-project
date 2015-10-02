package main

import (
	"fmt"
)

// multiply takes in two matrices and returns their product
func multiply(mat1, mat2 Matrix, c chan Matrix, i, j, a, b int) {
	dim1, dim2 := len(mat1.Val[0]), len(mat2.Val[0])
	if dim1 != dim2 {
		c <- Matrix{e: fmt.Errorf("Cannot multiply matrices of varying lengths: %v != %v", dim1, dim2)}
	} else if dim1 == 0 {
		c <- Matrix{e: fmt.Errorf("Empty matrices.")}
	}

	product := make([][]Element, dim1)

	for i := 0; i < dim1; i++ {
		product[i] = make([]Element, dim1)
		for j := 0; j < dim1; j++ {
			sum := 0
			for k := 0; k < dim1; k++ {
				sum += mat1.Val[i][k].Val * mat2.Val[k][j].Val
			}
			product[i][j] = Element{sum}
		}
	}

	c <- Matrix{product, nil, i, j, a, b}
}
