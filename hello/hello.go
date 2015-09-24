package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	// "strings"
)

// Creates a matrix of given input size and writes it to a file
func writeMatrix(width int, height int, fileName string) {
	fmt.Println("Start")
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	for i := 0; i < height; i++ {
		numList := ""
		for j := 0; j < width; j++ {
			numList += strconv.Itoa(rand.Intn(101)) + " "
		}
		file.WriteString(numList + "\n")
	}
}

func createMatrices(dim int, fileName1, fileName2 string) {
	go writeMatrix(dim, dim, fileName1)
	go writeMatrix(dim, dim, fileName2)
	fmt.Print("Success!")
}

func main() {
	createMatrices(1<<8, "input1", "input2")
}
