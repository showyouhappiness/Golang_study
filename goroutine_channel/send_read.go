package goroutine_channel

import "fmt"

func sendMessage(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func readMessage(ch <-chan string) {
	fmt.Printf("Receiving: %#v111111111\n", <-ch)
}

func SendReadMessage() {
	ch := make(chan string, 1)
	sendMessage(ch, "Hello World!")
	readMessage(ch)
}
