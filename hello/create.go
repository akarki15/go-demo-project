package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Creates a matrix of given input size and writes it to a file
func writeMatrix(width int, height int, fileName string, c chan string) {
	fmt.Println("Creating", fileName)
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
	c <- "done"
}

// Spawns two goroutines to create the two input files
func createMatrices(dim int, fileName1, fileName2 string) {
	c1 := make(chan string)
	c2 := make(chan string)

	go writeMatrix(dim, dim, fileName1, c1)
	go writeMatrix(dim, dim, fileName2, c2)

	go func() {
		for {
			select {
			case <-c1:
				fmt.Println("Done with", fileName1)
			case <-c2:
				fmt.Println("Done with", fileName2)
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
	createMatrices(1<<10, "input1", "input2")
}
