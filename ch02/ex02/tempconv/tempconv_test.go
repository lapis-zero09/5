package tempconv

import "fmt"

func Example_one() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F

	// Output:
	// 100
	// 180
}

func Example_two() {
	//!+printf
	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	//!-printf

	// Output:
	// 100°C
	// 100°C
	// 100°C
	// 100°C
}
