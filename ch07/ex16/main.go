package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/lapis-zero09/5/ch07/ex15/eval"
)

func parseInput(input string, env eval.Env) (eval.Expr, error) {
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

	for v := range vars {
		if _, ok := env[v]; !ok {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}

	return expr, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		print(w)
		r.ParseForm()
		env := eval.Env{}
		fmt.Println(r.Form)

		_, vs := r.Form["vars"]

		fmt.Println(vs)
		v, err := strconv.ParseFloat(vs[0], 64)
		if err != nil {
			fmt.Fprintf(w, "invalid variable %s=%q\n", k, vs[0])
			continue
		}
		env[eval.Var(k)] = v

		_, exprStr := r.Form["expr"]
		fmt.Fprint(w, exprStr, " => ")
		expr, err := parseInput(exprStr, env)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		fmt.Fprintln(w, expr.Eval(env))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	// http://localhost:8080/?expr=90%2B98323-32
}

func print(w io.Writer) {
	tmpl.Execute(w, nil)
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
		expr:<input type="text" name="expr">
		vars:<input type="text" name="vars">
    <input type="submit" value="">
	</form>
  </div>
</div>
</body>
</html>
`))
