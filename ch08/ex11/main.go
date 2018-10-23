package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func mirroredQuery(urls []string) []byte {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	responses := make(chan []byte)
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			b, err := request(ctx, u)
			if err != nil {
				log.Print(err)
				return
			}
			responses <- b
		}(url)
	}
	r := <-responses
	cancel()
	wg.Wait()
	return r
}

func request(ctx context.Context, url string) ([]byte, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func main() {
	b := mirroredQuery(os.Args[1:])
	fmt.Println(string(b))
}
