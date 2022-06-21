package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	//和python一样
	a, b = b, a
	t.Log(a, b)
}
