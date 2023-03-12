package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func hitURL(url string, c chan<- requestResult) { // send only
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}

type requestResult struct {
	url    string
	status string
}

func run() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	// url goroutines 조회
	for _, url := range urls {
		go hitURL(url, c)
	}
	// channel 메세지 받아오기
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	// 결과 출력
	for url, status := range results {
		fmt.Println(url, status)
	}
}
