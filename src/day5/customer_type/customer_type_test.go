package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n * 2) //只能写在打印这里，可以查看用时，否则得到的打印时长都为0
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op * op
}

func TestFn(t *testing.T) {
	tsSf := timeSpent(slowFunc)
	t.Log(tsSf(10))
}
