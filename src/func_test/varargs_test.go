package func_test

import "testing"

func F1(s ...string) {
	print(s)
	F2(s...)
	F3(s)
}

func F2(s ...string) {
	println(s)
}
func F3(s []string) {
	println(s)
}

func TestFunc(t *testing.T) {
	F1("a", "b", "c")
	Greeting("hello:", "Joe", "Anna", "Eileen")
}

func Greeting(prefix string, who ...string) {
	println(prefix, who)
}
