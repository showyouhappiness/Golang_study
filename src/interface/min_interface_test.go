/*
Sorter 接口，创建一个 Miner 接口并实现一些必要的操作。函数 Min() 接受一个 Miner 类型变量的集合，然后计算并返回集合中最小的元素。
*/
package interface_test

import (
	"fmt"
	"testing"
)

type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(i, j int) bool
	Swap(i, j int)
}

func Min(data Miner) interface{} {
	min := data.ElemIx(0)
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			min = data.ElemIx(i)
		} else {
			data.Swap(i, i-1)
		}
	}
	return min
}

func (p IntArray) ElemIx(ix int) interface{} { return p[ix] }

func (p StringArray) ElemIx(ix int) interface{} { return p[ix] }

func ints1() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data) //conversion to type IntArray
	m := Min(a)
	fmt.Printf("The minimum of the array is: %v\n", m)
}

func strings1() {
	data := []string{"ddd", "eee", "bbb", "ccc", "aaa"}
	a := StringArray(data)
	m := Min(a)
	fmt.Printf("The minimum of the array is: %v\n", m)
}

func TestMinInterface(t *testing.T) {
	ints1()
	strings1()
}
