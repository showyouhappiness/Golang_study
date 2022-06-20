package forExample

import (
	"fmt"
	"testing"
)

func TestJumpStatement(t *testing.T) {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

LABEL2:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

LABEL3:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				goto LABEL3
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")

HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
