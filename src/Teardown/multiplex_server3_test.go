package Teardown

import (
	"fmt"
	"testing"
)

func (r *Request) String() string {
	return fmt.Sprintf("%d+%d=%d", r.a, r.b, <-r.replyc)
}

func TestMultiplexServer3(t *testing.T) {
	adder, quit := startServer(func(a, b int) int { return a + b })
	// make requests:
	req1 := &Request{3, 4, make(chan int)}
	req2 := &Request{150, 250, make(chan int)}
	// send requests on the service channel
	adder <- req1
	adder <- req2
	// ask for the results: ( method String() is called )
	fmt.Println(req1, req2)
	// shutdown server:
	quit <- true
	fmt.Print("done")
}
