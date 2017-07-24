package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

var urls = []string{
	"http://pulsoconf.co/",
	"http://golang.org/",
	"http://matt.aimonetti.net/",
	"http://www.google.com/robots.txt",
	"http://yahoo.com",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncHttpGets(urls []string) <-chan *HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			resp.Body.Close()
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}
	return ch
}

func main() {

	zap := append(urls, urls...)
	fmt.Printf("Length of array %d\n",len(zap))


	results := asyncHttpGets(zap)
	for _ = range urls {
		result := <-results

		data, _ := ioutil.ReadAll(result.response.Body)

		fmt.Printf("%s status: %s\n body %s\n", result.url, result.response.Status, data)
	}
}