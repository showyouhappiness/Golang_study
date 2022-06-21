package func_test

import (
	"fmt"
	"testing"
)

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func TestChangeOutside(t *testing.T) {
	n := 0
	reply := &n
	print(reply)
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
}
