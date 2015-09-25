package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
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

	go func() {
		for {
			select {
			case <-c1:
				fmt.Println("Done with", filename1)
			case <-c2:
				fmt.Println("Done with", filename2)
			default:
				fmt.Print(".")
				time.Sleep(time.Second * 3)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)

	fmt.Print("Stopped!")
}

func main() {
	// createMatrices(4, "input1", "input2")
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
	m := make(map[string]Matrix)

	// read up parts of first multiplicand matrix
	for name, filename := range map[string]string{"a": filename1, "b": filename2} {
		var err error
		m[name+"11"], err = readSubMatrix(filename, 0, 0, dim/2)
		m[name+"12"], err = readSubMatrix(filename, 0, dim/2, dim/2)
		m[name+"21"], err = readSubMatrix(filename, dim/2, 0, dim/2)
		m[name+"22"], err = readSubMatrix(filename, dim/2, dim/2, dim/2)
		if err != nil {
			fmt.Println(err)
		}
	}
	for i, v := range m {
		fmt.Println(i, v)
	}
}
