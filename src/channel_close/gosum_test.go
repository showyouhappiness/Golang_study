//开启一个协程来计算 2 个整数的和并等待计算结果并打印出来。
package channel_close

import (
	"fmt"
	"testing"
)

func sum(x, y int, c chan int) {
	c <- x + y
}

func TestGoSum(t *testing.T) {
	c := make(chan int)
	go sum(12, 13, c)
	fmt.Println(<-c) // 25
}
