package main

import "fmt"

type A struct {
	B
	//name string
}

type B struct {
	name string
}

func main() {
	//a := A{B{name: "b"}, "a"}
	a := A{B{name: "b"}}
	fmt.Println(a.name, a.B.name)
}
