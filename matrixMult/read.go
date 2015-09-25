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
		// Get all ints as [] Element

		elementList := toElementList(strings.Split(scanner.Text(), " "))
		elementListLength := len(elementList)

		// declare the capacity of the slice
		mat = make([][]Element, elementListLength)
		mat[0] = elementList

		// loop through the rest
		count := 1
		for scanner.Scan() {
			elementList = toElementList(strings.Split(scanner.Text(), " "))
			mat[count] = elementList
			count++
		}
		return mat, nil
	}

	return nil, nil
}

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
