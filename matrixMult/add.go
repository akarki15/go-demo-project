package main

// // add takes in two matrices and returns their sum
// func add(mat1, mat2 Matrix, c chan Matrix) {
// 	dim := len(mat1[0])
// 	if dim == 0 {
// 		return
// 	}

// 	sum := make([][]Element, dim)
// 	for i := 0; i < dim; i++ {
// 		sum[i] = make([]Element, dim)
// 		for j := 0; j < dim; j++ {
// 			sum[i][j] = Element{mat1[i][j].Val + mat2[i][j].Val}
// 		}
// 	}
// 	c <- sum
// }
