package main

import (
	"fmt"

	"github.com/lapis-zero09/5/ch02/ex02/lengthconv"
	"github.com/lapis-zero09/5/ch02/ex02/tempconv"
	"github.com/lapis-zero09/5/ch02/ex02/weightconv"
)

func main() {
	fmt.Printf("Brrrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))

	fmt.Println("--------lengthconv--------")
	fmt.Printf("%s\n", lengthconv.FToM(123))
	fmt.Printf("%s\n", lengthconv.MToF(12))
	fmt.Printf("%s\n", lengthconv.FToM(0))
	fmt.Printf("%s\n", lengthconv.MToF(0))

	fmt.Println("--------weightconv--------")
	fmt.Printf("%s\n", weightconv.PToKg(123))
	fmt.Printf("%s\n", weightconv.KgToP(12))
	fmt.Printf("%s\n", weightconv.PToKg(0))
	fmt.Printf("%s\n", weightconv.KgToP(0))
}
