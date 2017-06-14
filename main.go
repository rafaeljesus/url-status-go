package main

import (
	"fmt"
	"net/http"
	"time"
)

type httpResponse struct {
	url    string
	status string
	err    error
}

func main() {
	ch := make(chan httpResponse)
	urls := [...]string{
		"https://github.com",
		"https://golang.org",
		"https://in.va.lid.zet",
	}

	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			res, err := http.Get(url)
			if err != nil {
				ch <- httpResponse{url: url, err: err}
				return
			}
			ch <- httpResponse{url, res.Status, err}
		}(url)
	}

	var count int
loop:
	for {
		select {
		case r := <-ch:
			count++
			if r.err != nil {
				fmt.Printf("%s -> %v \n", r.url, r.err)
				return
			}
			fmt.Printf("%s -> %v \n", r.url, r.status)
			if count == len(urls) {
				break loop
			}
		case <-time.After(5 * time.Second):
			count++
			fmt.Printf("timeout during fetching url \n")
			if count == len(urls) {
				break loop
			}
		}
	}
}
