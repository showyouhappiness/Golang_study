package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) { // 斐波那契数列
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

func TestExchange(t *testing.T) { // 交换两个数
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	//和python一样
	a, b = b, a
	t.Log(a, b)
}

func TestRecursive(t *testing.T) { // 递归 从10到1
	reduce(10)
}
func reduce(a int) {
	if a > 0 {
		print(a)
		a--
		reduce(a)
	}
}

func TestFactorial(t *testing.T) { // 递归 阶乘
	i := 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
