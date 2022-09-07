package http_web_server

import (
	"fmt"
	"net/http"
	"testing"
)

var urls = []string{
	"http://www.bing.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

func TestPollUrl(t *testing.T) {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
