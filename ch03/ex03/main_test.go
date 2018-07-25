package main

func ExampleCheckInf() {
	// checkInfAndPrint([]float64{0, 0, 0, 0, 0, 0, 0, math.NaN()}, 255, 255)
	// checkInfAndPrint([]float64{0, 0, 0, 0, 0, 0, 0, math.Inf(0)}, 255, 255)
	checkInfAndPrint([]float64{2.4, 2.4, 2.4, 2.4, 2.4, 2.4, 2.4, 2.4}, 255, 255)

	// Output:
	// <polygon points='2.4,2.4 2.4,2.4 2.4,2.4 2.4,2.4' fill='rgba(255,0,255,0.8)' stloke='rgba(255,0,255,0.8)'/>
}
