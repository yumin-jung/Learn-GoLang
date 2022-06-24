package main

import (
	"fmt"
	"net/http"
)

type reqRes struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)

	c := make(chan reqRes)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- reqRes) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- reqRes{url: url, status: "FAILED"}
	} else {
		c <- reqRes{url: url, status: "SUCCEED"}
	}
}
