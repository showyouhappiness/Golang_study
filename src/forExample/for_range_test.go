package forExample

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	numbers1 := [...]int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
}
