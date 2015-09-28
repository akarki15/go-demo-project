package main

// multiply takes in two matrices and returns their product
func multiply(mat1, mat2 Matrix, c chan Matrix) {
	dim := len(mat1[0])
	if dim == 0 {
		return
	}

	product := make([][]Element, dim)

	for i := 0; i < dim; i++ {
		product[i] = make([]Element, dim)
		for j := 0; j < dim; j++ {
			sum := 0
			for k := 0; k < dim; k++ {
				sum += mat1[i][k].Val * mat2[k][j].Val
			}
			product[i][j] = Element{sum}
		}
	}
	c <- product
}
