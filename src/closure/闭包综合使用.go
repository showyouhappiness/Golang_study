package main

import "fmt"

func func1() (i int) {
	i = 100
	defer func() {
		i += 1
	}()
	return 5
}

func func2() int {
	var i int
	defer func() {
		i += 1

	}()
	i += 100

	return i
}

func main() {
	fmt.Println(func1())
	fmt.Println(func2())
}
