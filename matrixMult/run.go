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
	Val [][]Element
	e   error

	// store metadata when used for multiplication and reading
	i, j int
	a, b int
}

func (m Matrix) String() string {
	st := "\n"
	if m.Val != nil {
		for i := 0; i < len(m.Val[0]); i++ {
			for j := 0; j < len(m.Val[0]); j++ {
				st = st + " " + strconv.Itoa(m.Val[i][j].Val)
			}
			st = st + "\n"
		}
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

	if err := createMatrices(dim, "input1", "input2"); err != nil {
		fmt.Println(err)
		return
	}

	if err := parallelMultiply("input1", "input2", dim); err != nil {
		fmt.Println(err)
		return
	}

}

func parallelMultiply(filename1, filename2 string, dim int) error {

	// channels to store temp data
	c1, c2 := make(chan Matrix), make(chan Matrix)

	// store input quadrants
	input1 := [2][2]Matrix{}
	input2 := [2][2]Matrix{}

	// read up parts of first multiplicand matrix
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			actualI, actualJ := (i * dim / 2), (j * dim / 2)
			go readSubMatrix(filename1, actualI, actualJ, dim/2, c1)
			go readSubMatrix(filename2, actualI, actualJ, dim/2, c2)
		}
	}

	// store input data since we are going to reuse some of them
	for i := 0; i < 4; i++ {
		temp1, temp2 := <-c1, <-c2
		input1[temp1.i][temp1.j], input2[temp2.i][temp2.j] = temp1, temp2
	}

	// multiply the quadrants
	alt := true
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				var c chan Matrix
				if alt {
					c = c1
				} else {
					c = c2
				}
				go multiply(input1[i][k], input2[k][j], c, i, k, k, j)
				alt = !alt
			}
		}
	}

	// store the products
	p1, p2 := [2][2]Matrix{}, [2][2]Matrix{}
	for i := 0; i < 4; i++ {
		temp1, temp2 := <-c1, <-c2
		p1[temp1.i][temp1.b] = temp1
		p2[temp2.i][temp2.b] = temp2
	}

	// add the products to get each quadrant of the final result sum
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			go add(p1[i][j], p2[i][j], c1)
			fmt.Println("Quadrant", i, j, <-c1)
		}
	}
	return nil
}
