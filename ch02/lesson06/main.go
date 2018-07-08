package main

import (
	"fmt"

	"github.com/lapis-zero09/5/ch02/lesson06/tempconv"
)

func main() {
	fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
