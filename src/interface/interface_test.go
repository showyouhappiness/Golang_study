package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (gp *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Logf(p.WriteHelloWorld())
}

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func TestShaper(t *testing.T) {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func TestInterfacesPoly(t *testing.T) {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}

//我的梦在哪里
//我就漂在哪里

//如果你们有梦想
//放开的去追

//因为只有奋斗
//才能改变命运
func TestInterfacesPoly2(t *testing.T) {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	c := &Circle{2.5}
	fmt.Println("Looping through shapes for area ...")
	// shapes := []Shaper{Shaper(r), Shaper(q), Shaper(c)}
	shapes := []Shaper{r, q, c}
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
