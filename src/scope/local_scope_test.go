package scope

import "testing"

var a = "G"

func TestLocalScope(t *testing.T) {
	n()
	m()
	n()
}

func n() { print(a, &a) }

func m() {
	a := "O" // 重新定义变量a，此时地址发生变化
	print(a, &a)
}
