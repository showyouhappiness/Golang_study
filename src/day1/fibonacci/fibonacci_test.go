package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	//和python一样
	a, b = b, a
	t.Log(a, b)
}
