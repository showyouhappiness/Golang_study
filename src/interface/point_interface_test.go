/*
定义接口 Magnitude，它有一个方法 Abs()。让 Point、Point3 及 Polar 实现此接口。通过接口类型变量使用方法做 point.go 中同样的事情。
*/
package interface_test

import (
	"fmt"
	"math"
	"testing"
)

type Magnitude interface {
	Abs() float64
}

var m Magnitude

type Point struct {
	X, Y float64
}

func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

type Point3 struct {
	X, Y, Z float64
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

type Polar struct {
	R, Theta float64
}

func (p Polar) Abs() float64 { return p.R }

func TestPointInterface(t *testing.T) {
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	m = p1 // p1 is type *Point, has method Abs()
	fmt.Printf("The length of the vector p1 is: %f\n", m.Abs())

	p2 := &Point{4, 5}
	m = p2
	fmt.Printf("The length of the vector p2 is: %f\n", m.Abs())

	p1.Scale(5)
	m = p1
	fmt.Printf("The length of the vector p1 after scaling is: %f\n", m.Abs())
	fmt.Printf("Point p1 after scaling has the following coordinates: X %f - Y %f\n", p1.X, p1.Y)

	mag := m.Abs()
	m = &Point3{3, 4, 5}
	mag += m.Abs()
	m = Polar{2.0, math.Pi / 2}
	mag += m.Abs()
	fmt.Printf("The float64 mag is now: %f", mag)
}
