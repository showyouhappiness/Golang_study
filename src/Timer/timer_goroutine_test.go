package Timer

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(1e8)
		}
	}
}
