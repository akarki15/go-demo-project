package main

import (
	// "errors"
	"fmt"
	"strconv"
)

type Element struct {
	Val int
}

type Matrix struct {
	Val  [][]Element
	e    error
	i, j int
}

func (m Matrix) String() string {
	st := "\n"
	for i := 0; i < len(m.Val[0]); i++ {
		for j := 0; j < len(m.Val[0]); j++ {
			st = st + " " + strconv.Itoa(m.Val[i][j].Val)
		}
		st = st + "\n"
	}
	return st
}

func (m1 Matrix) IsEqual(m2 Matrix) (bool, error) {
	dim1, dim2 := len(m1.Val[0]), len(m2.Val[0])
	if dim1 != dim2 {
		return false, fmt.Errorf("Cannot multiply matrices of varying lengths: %v != %v", dim1, dim2)
	}

	eq := true
	for i := 0; i < dim1; i++ {
		for j := 0; j < dim2; j++ {
			eq = eq && m1.Val[i][j] == m2.Val[i][j]
		}
	}
	return eq, nil
}

func main() {

	var dim int
	_, err := fmt.Scanf("%d", &dim)
	if err != nil {
		fmt.Println(err)
		return
	}

	createMatrices(dim, "input1", "input2")
	c := make(chan Matrix)
	d := make(chan Matrix)

	go readSubMatrix("input1", 0, 0, dim, c)
	go readSubMatrix("input2", 0, 0, dim, c)
	go multiply(<-c, <-c, d)
	prod := <-d
	if prod.e != nil {
		fmt.Println(prod.e)
	} else {
		fmt.Println(prod)
	}

	// parallelMultiply("input1", "input2", dim)

}

// func parallelMultiply(filename1, filename2 string, dim int) {
// 	fmt.Println(dim)
// 	// channels to store input quadrants
// 	m1 := [2][2]chan Matrix{}
// 	m2 := [2][2]chan Matrix{}

// 	// store input quadrants
// 	m1Mat := [2][2]Matrix{}
// 	m2Mat := [2][2]Matrix{}

// 	// channels to store products
// 	product1 := [2][2]chan Matrix{}
// 	product2 := [2][2]chan Matrix{}

// 	// channels to store sum
// 	sum := [2][2]chan Matrix{}

// 	// initialize
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 2; j++ {
// 			m1[i][j] = make(chan Matrix)
// 			m2[i][j] = make(chan Matrix)

// 			product1[i][j] = make(chan Matrix)
// 			product2[i][j] = make(chan Matrix)

// 			sum[i][j] = make(chan Matrix)
// 		}
// 	}

// 	// read up parts of first multiplicand matrix
// 	go readSubMatrix(filename1, 0, 0, dim/2, m1[0][0])
// 	go readSubMatrix(filename1, 0, dim/2, dim/2, m1[0][1])
// 	go readSubMatrix(filename1, dim/2, 0, dim/2, m1[1][0])
// 	go readSubMatrix(filename1, dim/2, dim/2, dim/2, m1[1][1])

// 	// read up parts of second multiplicand
// 	go readSubMatrix(filename2, 0, 0, dim/2, m2[0][0])
// 	go readSubMatrix(filename2, 0, dim/2, dim/2, m2[0][1])
// 	go readSubMatrix(filename2, dim/2, 0, dim/2, m2[1][0])
// 	go readSubMatrix(filename2, dim/2, dim/2, dim/2, m2[1][1])

// 	// store the received input matrix channel data
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 2; j++ {
// 			m1Mat[i][j] = <-m1[i][j]
// 			m2Mat[i][j] = <-m2[i][j]
// 		}
// 	}

// 	// multiply the quadrants
// 	alt := true
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 2; j++ {
// 			for k := 0; k < 2; k++ {
// 				if alt {
// 					go multiply(m1Mat[i][k], m2Mat[k][j], product1[i][j])
// 				} else {
// 					go multiply(m1Mat[i][k], m2Mat[k][j], product2[i][j])
// 				}
// 				alt = !alt
// 			}
// 		}
// 	}

// 	// add the products to get each quadrant of the final result sum
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 2; j++ {
// 			go add(<-product1[i][j], <-product2[i][j], sum[i][j])
// 			fmt.Println("Quadrant", i, j, <-sum[i][j])
// 		}
// 	}

// }
