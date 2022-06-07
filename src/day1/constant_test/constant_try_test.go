package constant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
)
const (
	// 位掩码表达式
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111  true true true
	//a := 1 //0001  true false false
	fmt.Println(Readable, Writable, Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
