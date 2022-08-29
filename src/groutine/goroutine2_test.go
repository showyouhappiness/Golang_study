package groutine

import (
	"fmt"
	"testing"
)

func TestGoroutine5(t *testing.T) {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)

		//for input := range ch {
		//	fmt.Printf("%s ", input)
		//}
	}
}
