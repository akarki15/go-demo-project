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
			elementList += strconv.Itoa(rand.Intn(11)) + " "
		}
		// last int
		elementList += strconv.Itoa(rand.Intn(11)) + "\n"
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
