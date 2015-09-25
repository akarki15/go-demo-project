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
	// createMatrices(10, "input1", "input2")
	fmt.Println(readMatrix("input1"))
}
