package main

import "fmt"

func rotate(x []int, r int) {
	idx := len(x) - r
	tmp := append(x[idx:], x[:idx]...)
	for i := 0; i < len(x); i++ {
		x[i] = tmp[i]
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(a)

	rotate(a, 3)
	fmt.Println(a)
}
