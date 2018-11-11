package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type PackageInfo struct {
	ImportPath string
	Deps       []string
}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("empty args")
		os.Exit(1)
	}

	pkgs, err := goList(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, pkg := range pkgs {
		ps, err := goList(pkg.Deps)
		// fmt.Println(ps)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(pkg.ImportPath)
		for _, p := range ps {
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(p.ImportPath)
		}
	}
}

func goList(pkgs []string) ([]PackageInfo, error) {
	args := []string{"list", "-e", "-json"}
	args = append(args, pkgs...)
	cmd := exec.Command("go", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	go cmd.Wait()
	dec := json.NewDecoder(stdout)
	var pis []PackageInfo
	for {
		p := PackageInfo{}
		err := dec.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		pis = append(pis, p)
	}
	return pis, nil
}
