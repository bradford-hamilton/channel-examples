package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// main will print whichever site responded the fastest
func main() {
	fastestResponse := multipleQuery()
	fmt.Println(fastestResponse)
}

func multipleQuery() string {
	// use buffered channel so we can make capacity of 3 and pull of the first
	responses := make(chan string, 3)

	go func() { responses <- request("http://google.com") }()
	go func() { responses <- request("http://reddit.com") }()
	go func() { responses <- request("https://news.ycombinator.com/news") }()

	return <-responses // return quickest response
}

func request(hostname string) (response string) {
	res, err := http.Get(hostname)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	bodyString := string(bodyBytes)

	return bodyString
}
