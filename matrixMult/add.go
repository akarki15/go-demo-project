package main

import (
	"fmt"
)
// add takes in two matrices and returns their sum
func add(mat1, mat2 Matrix, c chan Matrix) (chan Matrix, error) {	
	
	dim1, dim2 := len(mat1[0]), len(mat1[0])

	if dim1 != dim2 {
		return nil, fmt.Errorf("Cannot add matrices of different dimensions: %v != %v", dim1, dim2)
	} else if dim1 == 0 {
		return nil, nil
	}

	sum := make([][]Element, dim1)
	for i := 0; i < dim1; i++ {
		sum[i] = make([]Element, dim)
		for j := 0; j < dim1; j++ {
			sum[i][j] = Element{mat1[i][j].Val + mat2[i][j].Val}
		}
	}
	return c <- sum, nil
}
