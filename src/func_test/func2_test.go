package func_test

import "testing"

var z1, z2, z3 int

func TestDoSomething(t *testing.T) {
	x := 1
	y := 2
	z1, z2, z3 = test(x, y)
	print(z1, z2, z3)
	z2, z2, z3 = test1(x, y)
	print(z1, z2, z3)
	z2, z2, z3 = test2(x, y)
	print(z1, z2, z3)
}

func test(x int, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func test1(x int, y int) (int, int, int) {
	z1 = x + y
	z2 = x - y
	z3 = x * y
	return z1, z2, z3
}

func test2(x int, y int) (z1 int, z2 int, z3 int) {
	z1 = x + y
	z2 = x - y
	z3 = x * y
	return
}
