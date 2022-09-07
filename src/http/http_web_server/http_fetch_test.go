package http_web_server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestFetch(t *testing.T) {
	res, err := http.Get("http://www.bing.com")
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
