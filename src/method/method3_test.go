package method

import (
	"fmt"
	"math"
	"testing"
)

/*
展示了内嵌结构体上的方法可以直接在外层类型的实例上调用
*/
type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

func TestMethod(t *testing.T) {
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs()) // 打印 5
}

func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100.
}
