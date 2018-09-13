package main

import (
	"fmt"
	"sort"
)

func equals(s sort.Interface, i, j int) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !equals(s, i, j) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println([]int{1, 2, 3, 2, 1})
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1, 2, 3, 2, 1})))

	fmt.Println([]int{1, 2, 3, 4, 1})
	fmt.Println(IsPalindrome(sort.IntSlice([]int{1, 2, 3, 4, 1})))

}
