package wiki

import (
	"fmt"
	"net/http"
	"testing"
)

func viewHandlerPart(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, _ := load(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func TestWikiPart(t *testing.T) {
	http.HandleFunc("/view/", viewHandlerPart)
	http.ListenAndServe(":8080", nil)
}
