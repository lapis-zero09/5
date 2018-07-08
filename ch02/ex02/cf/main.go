package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/lapis-zero09/5/ch02/ex02/lengthconv"
	"github.com/lapis-zero09/5/ch02/ex02/weightconv"
	"github.com/lapis-zero09/5/ch02/lesson06/tempconv"
)

func conv(t float64) {
	// tempconv
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Println("--------temp---------")
	fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
	fmt.Printf("%s = %s\n", c, tempconv.CToF(c))

	// lengthconv
	ft := lengthconv.Feet(t)
	m := lengthconv.Meter(t)
	fmt.Println("--------length---------")
	fmt.Printf("%s = %s\n", ft, lengthconv.FToM(ft))
	fmt.Printf("%s = %s\n", m, lengthconv.MToF(m))

	// weightconv
	lb := weightconv.Pound(t)
	kg := weightconv.Kilogram(t)
	fmt.Println("--------weight---------")
	fmt.Printf("%s = %s\n", lb, weightconv.PToKg(lb))
	fmt.Printf("%s = %s\n", kg, weightconv.KgToP(kg))
}

func parseFloat(str string) float64 {
	t, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	return t
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t := parseFloat(arg)
			conv(t)
		}
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := parseFloat(input.Text())
		conv(t)
	}
}
