package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
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
	//a, b := returnMultiValues()
	//t.Log(a, b)
	a, _ := returnMultiValues()
	t.Log(a)
	tsSf := timeSpent(slowFunc)
	t.Log(tsSf(10))
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	defer Clear() //一般情况下会在最后完成；报错后也可以正常使用，可以释放资源、释放锁。
	t.Log("Start")
	panic("err") //异常中断
}
