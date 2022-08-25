//有两个协程，第一个提供数字 0，10，20，...，90 并将他们放入通道，第二个协程从通道中读取并打印。main() 等待两个协程完成后再结束。
package channel_close

import (
	"fmt"
	"testing"
)

// integer producer:
func numGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out)
}

// integer consumer:
func numEchoRange(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func TestProducerConsumer(t *testing.T) {
	numChan := make(chan int)
	done := make(chan bool)
	go numGen(0, 10, numChan)
	go numEchoRange(numChan, done)

	<-done
}
