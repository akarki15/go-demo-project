package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Element struct {
	Val int
}

type Matrix [][]Element

// readMatrix reads a file and returns a corresponding Matrix
func readMatrix(filename string) (Matrix, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Find the size of each row and declare the Matrix accordingly
	var mat Matrix

	if scanner.Scan() {
		// Get all ints as [] string

		valList := strings.Split(scanner.Text(), " ")
		valListLength := len(valList)
		for i := 0; i < valListLength; i++ {
			fmt.Println(i, valList[i])
		}
		// declare the capacity of the slice
		mat = make([][]Element, valListLength)
		mat[0] = make([]Element, valListLength)

		for i := 0; i < valListLength; i++ {
			valInt, err := strconv.Atoi(valList[i])
			if err != nil {
				return nil, err
			}
			// fmt.Println(i, valListLength, valInt)
			mat[0][i] = Element{valInt}
		}
	}

	// for scanner.Scan() {
	// 	if len(m2) == 0{

	// 	}
	// 	fmt.Println(scanner.Text(), " <<<<<<<<<<<<")
	// }
	return nil, nil
}
