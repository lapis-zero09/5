package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lapis-zero09/5/ch07/ex15/eval"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input expr:")
	input.Scan()
	expr, vars, err := parseInput(input.Text())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	env := eval.Env{}
	for v := range vars {
		fmt.Println("Input Variable")
		for {
			fmt.Printf("%s = ", v)
			input.Scan()
			val, err := strconv.ParseFloat(input.Text(), 64)
			if err == nil {
				env[v] = val
				break
			}
			fmt.Println(err)
			fmt.Println("error has occurred. Input again.")
		}
	}

	fmt.Printf("\nInput expr: %s\n", expr.String())
	fmt.Println("Result:", expr.Eval(env))
}

func parseInput(input string) (eval.Expr, map[eval.Var]bool, error) {
	if len(input) <= 0 {
		return nil, nil, fmt.Errorf("input expr is empty.")
	}

	expr, err := eval.Parse(input)
	if err != nil {
		return nil, nil, err
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, nil, err
	}

	return expr, vars, nil
}
