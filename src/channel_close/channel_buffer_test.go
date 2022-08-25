package channel_close

import (
	"fmt"
	"testing"
	"time"
)

type Empty interface{}
type semaphore chan Empty

func TestChannelBuffer(t *testing.T) {
	c := make(chan int, 50)
	go func() {
		time.Sleep(1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 20
	fmt.Println("sent", 10)

	data := [...]int{1, 2, 3, 4, 5, 6}
	for i, v := range data {
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}
}

/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}

// acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
