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
	// channels to store input quadrants
	m1 := [2][2]chan Matrix{}
	m2 := [2][2]chan Matrix{}

	// store input quadrants
	m1Mat := [2][2]Matrix{}
	m2Mat := [2][2]Matrix{}

	// channels to store products
	product1 := [2][2]chan Matrix{}
	product2 := [2][2]chan Matrix{}

	// channels to store sum
	sum := [2][2]chan Matrix{}

	// initialize
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			m1[i][j] = make(chan Matrix)
			m2[i][j] = make(chan Matrix)

			product1[i][j] = make(chan Matrix)
			product2[i][j] = make(chan Matrix)

			sum[i][j] = make(chan Matrix)
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

	// store the received input matrix channel data
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			m1Mat[i][j] = <-m1[i][j]
			m2Mat[i][j] = <-m2[i][j]
		}
	}

	// Multiply the quadrants
	alt := true
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				fmt.Println(i, k, "  *  ", k, j, " -> ", i, j, alt)

				if alt {
					go multiply(m1Mat[i][k], m2Mat[k][j], product1[i][j])
				} else {
					go multiply(m1Mat[i][k], m2Mat[k][j], product2[i][j])
				}
				alt = !alt
			}
		}
	}

	// add the products to get each quadrant of the final result sum
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			go add(<-product1[i][j], <-product2[i][j], sum[i][j])
			fmt.Println("Quadrant", i, j, <-sum[i][j])
		}
	}

}
