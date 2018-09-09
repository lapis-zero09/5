package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()
	panic("aaa")

}
