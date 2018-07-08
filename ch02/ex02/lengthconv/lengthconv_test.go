package lengthconv

import (
	"fmt"
)

func ExampleFToM() {
	meter := FToM(0)
	fmt.Printf("%g\n", meter)
	fmt.Println(meter.String())

	fmt.Println(FToM(123))
	// Output:
	// 0
	// 0m
	// 37.4904m
}

func ExampleMToF() {
	ft := MToF(0)
	fmt.Printf("%g\n", ft)
	fmt.Println(ft.String())

	fmt.Println(MToF(12))
	// Output:
	// 0
	// 0ft
	// 39.37007874015748ft
}
