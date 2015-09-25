package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// Creates a matrix of given input size and writes it to a file
func writeMatrix(width int, height int, filename string, c chan string) {
	fmt.Println("Creating", filename)
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	for i := 0; i < height; i++ {
		elementList := ""
		for j := 0; j < width-1; j++ {
			elementList += strconv.Itoa(rand.Intn(101)) + " "
		}
		// last int
		elementList += strconv.Itoa(rand.Intn(101)) + "\n"
		file.WriteString(elementList)
	}
	c <- "done"
}

// Spawns two goroutines to create the two input files
func createMatrices(dim int, filename1, filename2 string) {
	c1 := make(chan string)
	c2 := make(chan string)

	go writeMatrix(dim, dim, filename1, c1)
	go writeMatrix(dim, dim, filename2, c2)

	<-c1
	<-c2
	fmt.Print("Finished!")
}

func main() {
	// createMatrices(1<<10, "input1", "input2")
	// mat1, err := readSubMatrix("input1", 2, 0, 2)
	// mat2, err := readSubMatrix("input2", 0, 0, 16)

	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(mat1)
	// }
	parallelMultiply("input1", "input2", 4)
	// fmt.Println(multiply(mat1, mat2))
}

func parallelMultiply(filename1, filename2 string, dim int) {
	fmt.Println(dim)

	m1 := [2][2]chan Matrix{}
	m2 := [2][2]chan Matrix{}

	// initialize
	for i := 0; i < dim/2; i++ {
		for j := 0; j < dim/2; j++ {
			m1[i][j] = make(chan Matrix)
			m2[i][j] = make(chan Matrix)
		}
	}

	go readSubMatrix(filename1, 0, 0, dim/2, m1[0][0])
	// go readSubMatrix(filename1, 0, dim/1, dim/2, m1[0][1])

	fmt.Println("Test")
	temp := <-m1[0][0]
	fmt.Println("Test1")
	fmt.Println(temp)

	// read up parts of first multiplicand matrix

	go readSubMatrix(filename1, 0, 0, dim/2, m1[0][0])
	go readSubMatrix(filename1, 0, dim/2, dim/2, m1[0][1])
	go readSubMatrix(filename1, dim/2, 0, dim/2, m1[1][0])
	go readSubMatrix(filename1, dim/2, dim/2, dim/2, m1[1][1])

	go readSubMatrix(filename2, 0, 0, dim/2, m2[0][0])
	go readSubMatrix(filename2, 0, dim/2, dim/2, m2[0][1])
	go readSubMatrix(filename2, dim/2, 0, dim/2, m2[1][0])
	go readSubMatrix(filename2, dim/2, dim/2, dim/2, m2[1][1])

	for i := 0; i < dim/2; i++ {
		for j := 0; j < dim/2; j++ {
			fmt.Println(<-m1[i][j])
			fmt.Println(<-m2[i][j])
		}
	}
}
