package weightconv

import (
	"fmt"
)

func ExamplePToKg() {
	kg := PToKg(0)
	fmt.Printf("%g\n", kg)
	fmt.Println(kg.String())

	fmt.Println(PToKg(123))
	// Output:
	// 0
	// 0kg
	// 55.791861510000004kg
}

func ExampleKgToP() {
	lb := KgToP(0)
	fmt.Printf("%g\n", lb)
	fmt.Println(lb.String())

	fmt.Println(KgToP(12))
	// Output:
	// 0
	// 0lb
	// 26.455471462185308lb
}
