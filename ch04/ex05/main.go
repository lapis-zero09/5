package main

import "fmt"

func removeSameAdjacent(x []int) []int {
	if len(x) < 2 {
		return x
	}

	arr := []int{x[0]}
	for _, val := range x[1:] {
		if arr[len(arr)-1] != val {
			arr = append(arr, val)
		}
	}
	return arr
}

func main() {
	fmt.Println(removeSameAdjacent([]int{1, 2, 2, 2, 3, 5}))
}
