package unit_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func square(op int) int {
	return op * op
}

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d,the actual %d", inputs[i], expected[i], ret)
		}
	}
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
		//if ret != expected[i] {
		//	t.Errorf("input is %d, the expected is %d,the actual %d", inputs[i], expected[i], ret)
		//}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}
