package scope

import "testing"

var c string

func TestFunction(t *testing.T) {
	c = "G"
	print(c, &c)
	f1()
}

func f1() {
	c := "O"
	print(c, &c)
	f2()
}

func f2() {
	print(c, &c)
}
