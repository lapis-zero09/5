package main

import (
	"fmt"
	"os"

	"github.com/lapis-zero09/5/ch01/lesson02/echo"
)

func main() {
	fmt.Println("Echo1")
	ret := echo.Echo1(os.Args)
	fmt.Println(ret)

	fmt.Println("Echo2")
	ret = echo.Echo2(os.Args)
	fmt.Println(ret)

	fmt.Println("Echo3")
	ret = echo.Echo3(os.Args)
	fmt.Println(ret)
}
