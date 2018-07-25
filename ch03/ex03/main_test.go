package main

func ExampleCheckInf() {
	// checkInfAndPrint([]float64{0, 0, 0, 0, 0, 0, 0, math.NaN()})
	// checkInfAndPrint([]float64{0, 0, 0, 0, 0, 0, 0, math.Inf(0)})
	checkInfAndPrint([]float64{12.423, 12.423, 12.423, 12.423, 12.423, 12.423, 12.423, 12.423})

	// Output:
	// <polygon points='12.423,12.423 12.423,12.423 12.423,12.423 12.423,12.423'/>
}
