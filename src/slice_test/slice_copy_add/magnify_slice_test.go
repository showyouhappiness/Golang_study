/*
给定一个切片 s []int 和一个 int 类型的因子 factor，扩展 s 使其长度为 len(s) * factor。

用顺序函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。通过 fn 测试方法测试当整型值是偶数时的情况。
*/
package slice_copy_add

import (
	"fmt"
	"testing"
)

var s []int

func TestMagnifySlice(t *testing.T) {
	s = []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(s))
	fmt.Println(s)
	s = enlarge(s, 5)
	fmt.Println("The length of s after enlarging is:", len(s))
	fmt.Println(s)
}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	fmt.Println("The length of ns  is:", len(ns))
	copy(ns, s)
	fmt.Println(ns)
	s = ns
	fmt.Println(s)
	fmt.Println("The length of s after enlarging is:", len(s))
	return s
}
