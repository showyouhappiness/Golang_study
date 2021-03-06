/*
程序中定义了一个新类型 Circle，它也实现了 Shaper1 接口。 if t, ok := areaIntf.(*Square);
ok 测试 areaIntf 里是否有一个包含 *Square 类型的变量，结果是确定的；然后我们测试它是否包含一个 *Circle 类型的变量，结果是否定的。
如果忽略 areaIntf.(*Square) 中的 * 号，会导致编译错误：impossible type assertion: Square does not implement Shaper1 (Area method has pointer receiver)。
链接
*/
package interface_test

import (
	"fmt"
	"math"
	"testing"
)

type Square1 struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper1 interface {
	Area() float32
}

func TestTypeInterface(t *testing.T) {
	var areaIntf Shaper1
	sq1 := new(Square1)
	sq1.side = 5

	areaIntf = sq1
	// Is Square1 the type of areaIntf?
	if t, ok := areaIntf.(*Square1); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}

func (sq *Square1) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
