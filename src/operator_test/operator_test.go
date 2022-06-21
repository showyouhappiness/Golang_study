package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	b1 := [...]int{1, 2, 4, 3}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)  //元素不同
	t.Log(a == b1) //顺序不同
	//t.Log(a == c) //长度不同，不能进行比较
	t.Log(a == d)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111  true true true
	//使用位运算符改变状态
	a = a &^ Readable
	a = a &^ Writable
	a = a &^ Executable
	//a := 1 //0001  true false false
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
