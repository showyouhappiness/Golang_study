/*
用顺序函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。通过 fn 测试方法测试当整型值是偶数时的情况。
*/
package slice_copy_add

import (
	"fmt"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = Filter(s, even)
	fmt.Println(s)
}

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, i := range s {
		if fn(i) {
			p = append(p, i)
		}
	}
	return p
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

/* [0 2 4 6 8] */
