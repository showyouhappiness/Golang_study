package http_web_server

import (
	"net/http"
	"testing"
)

func TestSimpleForm(t *testing.T) {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
