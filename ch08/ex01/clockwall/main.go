package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type clock struct {
	address string
	port    string
}

type time struct {
	idx  int
	text string
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	var clocks []*clock
	for _, server := range os.Args[1:] {
		clock, err := parseInput(server)
		if err != nil {
			return err
		}
		clocks = append(clocks, clock)
	}

	ch := make(chan time, len(clocks))
	for i, clock := range clocks {
		conn, err := net.Dial("tcp", clock.port)
		if err != nil {
			return err
		}

		fmt.Printf("%10s", clock.address)
		defer conn.Close()
		go handleConn(conn, i, ch)
	}
	fmt.Println()

	times := make([]string, len(clocks))
	for t := range ch {
		times[t.idx] = t.text
		fmt.Print("\r")
		for _, time := range times {
			fmt.Printf("%10s", time)
		}
	}
	return nil
}

func parseInput(input string) (*clock, error) {
	tokens := strings.Split(input, "=")
	if len(tokens) != 2 {
		return nil, fmt.Errorf("Invalid format address expcted: (address=localhost:port), but got (%s)", input)
	}

	return &clock{
		address: tokens[0],
		port:    tokens[1],
	}, nil
}

func handleConn(c net.Conn, i int, ch chan time) {
	b := bufio.NewScanner(c)
	for b.Scan() {
		ch <- time{i, b.Text()}
	}
}
