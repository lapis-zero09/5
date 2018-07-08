package main

import (
	"fmt"

	"github.com/lapis-zero09/5/ch02/ex01/tempconv"
)

func main() {
	fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.CToK(0))
}
