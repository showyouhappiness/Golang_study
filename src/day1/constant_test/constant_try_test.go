package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
)
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	//a := 7 //0111  true true true
	a := 1 //0001  true false false
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
