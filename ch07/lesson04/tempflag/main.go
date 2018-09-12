package main

import (
	"flag"
	"fmt"

	"github.com/lapis-zero09/5/ch07/lesson04/tempconv"
)

//!+
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
