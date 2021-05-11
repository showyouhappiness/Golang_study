package func_test

import (
	"fmt"
	"testing"
)

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	defer Clear() //一般情况下会在最后完成；报错后也可以正常使用，可以释放资源、释放锁。
	t.Log("Start")
	panic("err") //异常中断
}
