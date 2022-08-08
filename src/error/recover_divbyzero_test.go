package error

import (
	"fmt"
	"testing"
)

func badCall1() {
	a, b := 10, 0
	n := a / b
	fmt.Println(n)
}

func test1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}

	}()
	badCall1()
	fmt.Printf("After bad call\r\n")
}

func TestRecoverDivbyzero(t *testing.T) {
	fmt.Printf("Calling test\r\n")
	test1()
	fmt.Printf("Test completed\r\n")
}
