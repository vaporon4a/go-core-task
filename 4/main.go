package main

import "fmt"

// uniqueOnFirst takes two slices of strings and returns a new slice containing
// the elements that are present in the first slice but not in the second slice.
// It uses a map to store elements of the second slice for efficient lookup.

func uniqueOfFirst(slice1 []string, slice2 []string) []string {
	list := make(map[string]struct{}, len(slice1))

	for _, v := range slice2 {
		list[v] = struct{}{}
	}

	result := make([]string, 0, len(slice1))

	for _, v := range slice1 {
		if _, ok := list[v]; !ok {
			result = append(result, v)
			list[v] = struct{}{}
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(uniqueOfFirst(slice1, slice2))
}
