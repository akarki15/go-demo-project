package main

// import (
// 	"fmt"
// )

// // multiply takes in two matrices and returns their product
// func multiply(mat1, mat2 Matrix, c chan Matrix) err {
// 	dim1, dim2 := len(mat1[0]), len(mat2[0])
// 	if dim1 == dim2 {
// 		return fmt.Errorf("Cannot multiply matrices of varying lengths: %v != %v", dim1, dim2)
// 	} else if dim1 == 0 {
// 		return fmt.Errorf("Empty matrices.")
// 	}

// 	product := make([][]Element, dim1)

// 	for i := 0; i < dim1; i++ {
// 		product[i] = make([]Element, dim1)
// 		for j := 0; j < dim1; j++ {
// 			sum := 0
// 			for k := 0; k < dim1; k++ {
// 				sum += mat1[i][k].Val * mat2[k][j].Val
// 			}
// 			product[i][j] = Element{sum}
// 		}
// 	}

// 	c <- product
// 	return nil
// }
