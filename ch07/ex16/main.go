package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/lapis-zero09/5/ch07/ex15/eval"
)

func parseInput(input string) (eval.Expr, error) {
	if len(input) <= 0 {
		return nil, fmt.Errorf("input expr is empty.")
	}

	expr, err := eval.Parse(input)
	if err != nil {
		return nil, err
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}

	return expr, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
		r.ParseForm()

		exprStr, ok := r.Form["expr"]
		if !ok {
			return
		}
		fmt.Fprint(w, exprStr[0], " => ")
		expr, err := parseInput(exprStr[0])
		if err != nil {
			fmt.Fprintln(w, err)
		}
		fmt.Fprintln(w, expr.Eval(eval.Env{}))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var tmpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
	<title>calc</title>
	<link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
</head>
<body>
<div class="jumbotron jumbotron-fluid">
  <div class="container">
	<h1>calc</h1>
	<form action="/" method="post">
		expr: <input type="text" name="expr">
    <input type="submit" value="calc">
	</form>
  </div>
</div>
</body>
</html>
`))
