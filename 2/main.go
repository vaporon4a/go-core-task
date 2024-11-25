package main

import (
	"fmt"
	"math/rand"
)

const (
	sliceSize = 10
)

// sliceExample takes a slice of integers and returns a new slice containing only the even numbers.
func sliceExample(slice []int) []int {

	newSlice := make([]int, 0)

	for _, v := range slice {
		if v%2 == 0 {
			newSlice = addElements(newSlice, v)
		}
	}
	return newSlice
}

// copySlice creates a copy of the given slice of integers and returns it.
// It allocates a new slice with the same length as the input slice and copies each element.
func copySlice(slice []int) []int {

	newSlice := make([]int, len(slice))

	for i, v := range slice {
		newSlice[i] = v
	}
	return newSlice
}

// addElements takes a slice of integers and a number and returns a new slice containing all the elements of the original slice plus the given number.
//
// The new slice is allocated with a length one greater than the original and the elements of the original slice are copied into the new slice.
// The given number is then appended to the end of the new slice.
func addElements(slice []int, number int) []int {

	newSlice := make([]int, len(slice)+1)

	for i, v := range slice {
		newSlice[i] = v
	}
	newSlice[len(slice)] = number
	return newSlice
}

// removeElement removes the element at the specified index from the given slice.
// It returns a new slice with the element removed and an error if the index is out of range.
func removeElement(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return nil, fmt.Errorf("index out of range")
	}
	newSlice := make([]int, len(slice)-1)

	for i, v := range slice {
		switch {
		case i < index:
			newSlice[i] = v
		case i > index:
			newSlice[i-1] = v
		}
	}
	return newSlice, nil
}

func main() {
	originalSlice := make([]int, sliceSize)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(sliceSize)
	}
	fmt.Printf("Original slice:\t\t %v\n", originalSlice)
	fmt.Printf("Even slice:\t\t %v\n", sliceExample(originalSlice))

	newSlice := copySlice(originalSlice)
	fmt.Printf("Copied slice:\t\t %v\n", newSlice)
	remIndex := rand.Intn(len(newSlice))
	remSlice, err := removeElement(newSlice, remIndex)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Removed element[%d]:\t %v\n", remIndex, remSlice)
	fmt.Printf("Original slice:\t\t %v\n", originalSlice)
	fmt.Printf("Add element:\t\t %v\n", addElements(originalSlice, 5))
}
