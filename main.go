package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	status string
	body   []byte
}

func main() {
	resChan := make(chan Response)

	urls := [2]string{
		"https://github.com",
		"https://golang.org/",
	}

	for _, url := range urls {
		go checkUrl(url, resChan)
	}

	for range urls {
		fmt.Println((<-resChan).status)
	}
}

func checkUrl(url string, ch chan<- Response) {
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	ch <- Response{res.Status, body}
}
