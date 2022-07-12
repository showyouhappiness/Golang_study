/*
a). 继续扩展程序，定义类型 Triangle，让它实现 AreaInterface 接口。通过计算一个特定三角形的面积来进行测试（三角形面积=0.5 * (底 * 高)）
b). 定义一个新接口 PeriInterface，它有一个 Perimeter 方法。让 Square 实现这个接口，并通过一个 Square 示例来测试它。
*/
package interface_test

import (
	"fmt"
	"testing"
)

type Triangle struct {
	base   float32
	height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	Perimeter() float32
}

func TestInterfaceExt(t *testing.T) {
	var areaIntf AreaInterface
	var periIntf PeriInterface

	sq1 := new(Square)
	sq1.side = 5
	tr1 := new(Triangle)
	tr1.base = 3
	tr1.height = 5

	areaIntf = sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	periIntf = sq1
	fmt.Printf("The square has perimeter: %f\n", periIntf.Perimeter())

	areaIntf = tr1
	fmt.Printf("The triangle has area: %f\n", areaIntf.Area())
}

func (sq *Square) Perimeter() float32 {
	return 4 * sq.side
}

func (tr *Triangle) Area() float32 {
	return 0.5 * tr.base * tr.height
}
