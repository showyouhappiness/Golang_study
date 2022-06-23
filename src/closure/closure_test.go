package closure_test

import (
	"fmt"
	"strings"
	"testing"
)

func work() func() int {
	days := 0
	return func() int {
		days++
		return days
	}
}

func TestClosure(t *testing.T) {
	closure := work()
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
}

func func1() (i int) {
	i = 100
	defer func() {
		i += 1
	}()
	return 5
}

func func2() int {
	var i int
	defer func() {
		i += 1

	}()
	i += 100

	return i
}

func TestClosure1(t *testing.T) {
	fmt.Println(func1())
	fmt.Println(func2())
}

func TestClosure2(t *testing.T) {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func TestClosure3(t *testing.T) {
	var f = Adder1()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder1() func(int) int {
	var x int
	return func(delta int) int {
		println(delta)
		x += delta
		return x
	}
}
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func TestClosure4(t *testing.T) {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	println(addBmp("file"))
	println(addJpeg("file"))
}
