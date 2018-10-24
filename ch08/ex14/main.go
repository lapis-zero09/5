package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			cli.ch <- fmt.Sprintf("%d users are in this chat!", len(clients))
			for c := range clients {
				cli.ch <- c.name
			}
			cli.ch <- "Enjoy!!!"

			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	ch <- "Input your name!"
	input := bufio.NewScanner(conn)
	input.Scan()
	who := input.Text() + "@" + conn.RemoteAddr().String()

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	textCh := make(chan string)
	go func() {
		for input.Scan() {
			textCh <- who + ": " + input.Text()
		}
	}()

	defer func() {
		leaving <- client{ch, who}
		messages <- who + " has left"
		conn.Close()
	}()

	for {
		select {
		case message, ok := <-textCh:
			if !ok {
				return
			}
			messages <- message
		case <-time.After(5 * time.Minute):
			return
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
