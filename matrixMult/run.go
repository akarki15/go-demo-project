package main

import (
	"errors"
	"fmt"
)

type Element struct {
	Val int
}

type Matrix [][]Element

func (m1 Matrix) IsEqual(m2 Matrix) bool {
	if len(m1[0]) != len(m2[0]) {
		fmt.Println(errors.New("Can't compare matrices of different sizes."))
		return false
	}

	eq := true
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1[0]); j++ {
			eq = eq && m1[i][j] == m2[i][j]
		}
	}
	return eq
}

func main() {
	var dim int
	_, err := fmt.Scanf("%d", &dim)
	if err != nil {
		fmt.Println(err)
		return
	}
	createMatrices(dim, "input1", "input2")
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

	// multiply the quadrants
	alt := true
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				var dest Matrix
				if alt {
					dest = product1
				} else {
					dest = product2
				}
				go multiply(m1Mat[i][k], m2Mat[k][j], dest[i][j])
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
