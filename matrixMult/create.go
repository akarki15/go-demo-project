package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// Creates a matrix of given input size and writes it to a file
func writeMatrix(width int, height int, filename string, c chan error) {
	fmt.Println("Creating", filename)
	file, err := os.Create(filename)
	if err != nil {
		c <- err
	}
	defer file.Close()

	for i := 0; i < height; i++ {
		elementList := ""
		for j := 0; j < width-1; j++ {
			elementList += strconv.Itoa(rand.Intn(11)) + " "
		}
		// last int
		elementList += strconv.Itoa(rand.Intn(11)) + "\n"
		file.WriteString(elementList)
	}
	c <- nil
}

// Spawns two goroutines to create the two input files
func createMatrices(dim int, filename1, filename2 string) error {
	c := make(chan error)

	go writeMatrix(dim, dim, filename1, c)
	go writeMatrix(dim, dim, filename2, c)

	err1, err2 := <-c, <-c

	if err1 == nil && err2 == nil {
		return nil
	}

	if err1 != nil {
		fmt.Println(err1)
	}

	if err2 != nil {
		fmt.Println(err2)
	}
	return fmt.Errorf("Problem creating matrix.")
}
