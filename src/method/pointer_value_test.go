package method

import (
	"fmt"
	"math"
	"testing"
)

type B struct {
	thing int
}

type Point3 struct {
	x, y, z float64
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }
func (p Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func TestPointerValue(t *testing.T) {
	/*
		change()接受一个指向 B 的指针，并改变它内部的成员；
		write() 通过拷贝接受 B 的值并只输出 B 的内容。
		注意 Go 为我们做了探测工作，我们并没有指出是否在指针上调用方法，Go 替我们做了这些事情。b1 是值而 b2 是指针，方法都支持运行了。
	*/
	var b1 B // b1 是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2 是指针
	b2.change()
	fmt.Println(b2.write())

	/*
		在 write() 中改变接收者 b 的值：它可以正常编译，但是开始的 b 没有被改变。方法将指针作为接收者不是必须的
		这样做稍微有点浪费内存，因为 Point3 是作为值传递给方法的，因此传递的是它的拷贝，这在 Go 中是合法的。也可以在指向这个类型的指针上调用此方法（会自动解引用）。
	*/

	p := Point3{3, 4, 5}
	fmt.Println(p.Abs())

	p1 := new(Point3)
	p1.x = 3
	p1.y = 4
	p1.z = 5
	fmt.Println(p1.Abs())

	p2 := &Point3{3, 4, 5}
	fmt.Println(p2.Abs())
	fmt.Println((*p2).Abs())
}
