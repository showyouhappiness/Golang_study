package main

import "fmt"

func main() {
	var a interface{}
	fmt.Println(a == nil)

	var b *int = nil
	a = b
	fmt.Println(a == nil)
}
