package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	var results = make(map[string]string)

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
		result := "SUCCEED"
		err := hitURL(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
	}

	for url, res := range results {
		fmt.Println(res, url)
	}
}

var errReqFailed = errors.New("Request failed")

func hitURL(url string) error {
	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400 {
		return errReqFailed
	}
	return nil
}
