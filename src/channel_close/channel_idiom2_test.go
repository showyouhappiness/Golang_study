//给通道使用 for 循环从指定通道中读取数据直到通道关闭，才继续执行下边的代码。
package channel_close

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelIdiom2(t *testing.T) {
	suck1(pump())
	time.Sleep(1e9)
}

func suck1(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
