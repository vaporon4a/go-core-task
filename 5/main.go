package main

import "fmt"

func slicesCrossing(slice1 []int, slice2 []int) (bool, []int) {
	list := make(map[int]struct{}, len(slice1))
	var flag bool
	for _, v := range slice1 {
		list[v] = struct{}{}
	}
	capacity := max(len(slice1), len(slice2))
	result := make([]int, 0, capacity)
	for _, v := range slice2 {
		if _, ok := list[v]; ok {
			flag = true
			result = append(result, v)
			delete(list, v)
		}
	}

	return flag, result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	flag, result := slicesCrossing(a, b)
	fmt.Println(flag, result)
}
