package forExample

import (
	"fmt"
	"testing"
)

func TestForExample(t *testing.T) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			println(i, j)
		}
	}

	for i := 0; i < 15; i++ {
		println(i)
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < i; j++ {
			print("G")
		}
		print("\n")
	}

	// 1 - use 2 nested for loops
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	// 2 -  use only one for loop and string concatenation
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}

	Fizz := 3
	Buzz := 5
	FizzBuzz := 15
	for i := 0; i <= 100; i++ {
		switch {
		case i%Fizz == 0:
			fmt.Println("Fizz")
		case i%Buzz == 0:
			fmt.Println("Buzz")
		case i%FizzBuzz == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)
		}
	}
}
