package channel_close

import (
	"fmt"
	"testing"
)

func TestRandomBitgen(t *testing.T) {
	c := make(chan int)
	// consumer:
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()
	// producer:
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}

}
