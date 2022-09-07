package http_web_server

import (
	"fmt"
	"net/http"
	"testing"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func TestHelloServer(t *testing.T) {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
