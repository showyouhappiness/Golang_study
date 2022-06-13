package goroutine_channel

import (
	"fmt"
)

func send(ch chan string, message string) {
	ch <- message
}

func Channel2() {
	//size := 4
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	send(ch, "three")
	send(ch, "four")
	//go send(ch, "three")
	//go send(ch, "four")
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}
