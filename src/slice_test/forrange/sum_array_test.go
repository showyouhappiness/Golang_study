/*
写一个 Sum() 函数，传入参数为一个 float32 数组成的数组 arrF，返回该数组的所有数字和。
如果把数组修改为切片的话代码要做怎样的修改？如果用切片形式方法实现不同长度数组的的和呢？
写一个 SumAndAverage() 方法，返回两个 int 和 float32 类型的未命名变量的和与平均值。
*/
package forrange

import (
	"fmt"
	"testing"
)

func TestSumArray(t *testing.T) {
	var a = [4]float32{1.0, 2.0, 3.0, 4.0}
	//var a = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("The sum of the array is: %f\n", Sum(a))
	var b = []int{1, 2, 3, 4, 5}
	sum, average := SumAndAverage(b)
	fmt.Printf("The sum of the array is: %d, and the average is: %f", sum, average)
}

func Sum(a [4]float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}

//func Sum(a []float32) (sum float32) {
//	for _, item := range a {
//		sum += item
//	}
//	return
//}

func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum, float32(sum / len(a))
}
