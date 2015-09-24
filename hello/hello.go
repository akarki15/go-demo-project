package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func createMatrix(width int, height int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	for i := 0; i < length; i++ {
		for i := 0; i < height; i++ {
			file.WriteString(strings.Join(rand.Intn(101), " "))
		}
		file.WriteString("\n")
	}
	file.WriteString("test")
	fmt.Println("Success")
}

func main() {
	createMatrix(10, 10, "testFile.txt")
}
