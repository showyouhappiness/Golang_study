/*
一个 map-function 是指能够接受一个函数原型和一个列表，并使用列表中的值依次执行函数原型，公式为：map ( F(), (e1,e2, . . . ,en) ) = ( F(e1), F(e2), ... F(en) )。
编写一个函数 mapFunc 要求接受以下 2 个参数：
	一个将整数乘以 10 的函数
	一个整数列表
最后返回保存运行结果的整数列表。
*/
package string

import (
	"fmt"
	"testing"
)

func TestMapFunc(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7}
	mf := func(i int) int {
		return i * 10
	}
	/*
		result := mapFunc(mf, list)
		for _, v := range result {
			fmt.Println(v)
		}
	*/
	println()
	// shorter:
	fmt.Printf("%v", mapFunc(mf, list))
}

func mapFunc(mf func(int) int, list []int) []int {
	result := make([]int, len(list))
	for ix, v := range list {
		result[ix] = mf(v)
	}
	/*
		for ix := 0; ix<len(list); ix++ {
			result[ix] = mf(list[ix])
		}
	*/
	return result
}
