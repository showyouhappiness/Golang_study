package main

import "fmt"

func work() func() int {
	days := 0
	return func() int {
		days++
		return days
	}
}

func main() {
	closure := work()
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
}
