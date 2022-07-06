package method

import (
	"fmt"
	"testing"
)

type TwoInts struct {
	a int
	b int
}

type IntVector []int

/*
结构体上的简单方法的例子
*/
func TestStruct(t *testing.T) {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

/*
非结构体类型上方法的例子
*/

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func TestUnStruct(t *testing.T) {
	fmt.Println(IntVector{1, 2, 3}.Sum()) // 输出是6
}
