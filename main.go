package main

import (
	"fmt"
	"time"
)

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// )

// var errRequestFailed = errors.New("Request failed")

// func hitURL(url string) error {
// 	fmt.Println("Checking:", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println(err, resp.StatusCode)
// 		return errRequestFailed
// 	}
// 	return nil
// }

// func main() {
// 	// var results = map[string]string{}
// 	var results = make(map[string]string)
// 	urls := []string{
// 		"https://www.airbnb.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.reddit.com/",
// 		"https://www.google.com/",
// 		"https://soundcloud.com/",
// 		"https://www.facebook.com/",
// 		"https://www.instagram.com/",
// 		"https://academy.nomadcoders.co/",
// 	}
// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}
// 	// fmt.Println(results)
// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}
// }

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "lynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	// time.Sleep(time.Second * 5)
	result := <-c // 메세지 1개 받을때까지 기다림
	fmt.Println(result)
	fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
