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

// readMatrix reads a matrix file and returns matrix corresponding to the row, col and
// intermediateLen. For example if the matrix is of dimenstion 2*2 then, readSubMatrix
// (filename, 2, 2, 2) returns the bottom right quadrant

func readSubMatrix(filename string, row, col, intermediateLen int) (Matrix, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var mat Matrix = make([][]Element, intermediateLen)

	// ignore first 0 to row-1 records
	for i := 0; i < row; i++ {
		scanner.Scan()
	}

	count := 0
	for scanner.Scan() && count < intermediateLen {
		// get record [col:col+intermediateLen-1]
		elementList := toElementList(strings.Split(scanner.Text(), " ")[col : col+intermediateLen])
		fmt.Println(elementList)
		mat[count] = elementList
		count++
	}
	return mat, nil
}

// takes a list of nums of string type and returns a corresponding list of elements
func toElementList(stringList []string) []Element {

	elementList := make([]Element, len(stringList))

	for i := 0; i < len(stringList); i++ {
		val, err := strconv.Atoi(stringList[i])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		elementList[i] = Element{val}
	}
	return elementList
}
