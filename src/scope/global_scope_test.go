package scope

import "testing"

var b = "G"

func TestGlobalScope(t *testing.T) {
	x()
	y()
	x()
}

func x() {
	print(b, &b)
}

func y() {
	b = "O" // 给b赋值，此时地址不变
	print(b, &b)
}
