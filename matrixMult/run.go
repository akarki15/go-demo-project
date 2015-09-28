package main

import (
	"fmt"
)

type Element struct {
	Val int
}

type Matrix [][]Element

func main() {
	dim := 2
	createMatrices(dim, "input1", "input2")
	// mat1, err := readSubMatrix("input1", 2, 0, 2)
	// mat2, err := readSubMatrix("input2", 0, 0, 16)

	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(mat1)
	// }
	parallelMultiply("input1", "input2", dim)

}

func parallelMultiply(filename1, filename2 string, dim int) {
	fmt.Println(dim)

	m1 := [2][2]chan Matrix{}
	m2 := [2][2]chan Matrix{}

	product1 := [2][2]chan Matrix{}
	product2 := [2][2]chan Matrix{}

	// initialize
	for i := 0; i < dim/2; i++ {
		for j := 0; j < dim/2; j++ {
			m1[i][j] = make(chan Matrix)
			m2[i][j] = make(chan Matrix)
			product1[i][j] = make(chan Matrix)
			product2[i][j] = make(chan Matrix)
		}
	}

	// read up parts of first multiplicand matrix

	go readSubMatrix(filename1, 0, 0, dim/2, m1[0][0])
	go readSubMatrix(filename1, 0, dim/2, dim/2, m1[0][1])
	go readSubMatrix(filename1, dim/2, 0, dim/2, m1[1][0])
	go readSubMatrix(filename1, dim/2, dim/2, dim/2, m1[1][1])

	// read up parts of second multiplicand
	go readSubMatrix(filename2, 0, 0, dim/2, m2[0][0])
	go readSubMatrix(filename2, 0, dim/2, dim/2, m2[0][1])
	go readSubMatrix(filename2, dim/2, 0, dim/2, m2[1][0])
	go readSubMatrix(filename2, dim/2, dim/2, dim/2, m2[1][1])

	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Println(<-m1[i][j], <-m2[i][j])
	// 	}
	// }

	// Multiply parts
	alt := true

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				fmt.Println(i, k, "  *  ", k, j, " -> ", i, j, alt)
				fmt.Println(i, k, <-m1[i][k], <-m2[k][j], k, j)
				// // if alt {
				// // 	go multiply(<-m1[i][k], <-m2[k][j], product1[i][j])
				// // } else {
				// // 	go multiply(<-m1[i][k], <-m2[k][j], product2[i][j])
				// // }
				// alt = !alt
			}
		}
	}

	// fmt.Println("TEST")
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Println(<-product1[i][j])
	// 		fmt.Println(<-product2[i][j])
	// 	}
	// }

	// go multiply(<-m1[0][0], <-m2[0][0], product1[0][0])

	// fmt.Println(<-product[0][0])

	// for i := 0; i < dim/2; i++ {
	// 	for j := 0; j < dim/2; j++ {
	// 		fmt.Println(<-m1[i][j])
	// 		fmt.Println(<-m2[i][j])
	// 	}
	// }
}
