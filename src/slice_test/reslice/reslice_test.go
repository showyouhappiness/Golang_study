package reslice

import (
	"fmt"
	"testing"
)

func TestReSlice(t *testing.T) {
	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}
func TestReSliceInt(t *testing.T) {
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(ar)
	var a = ar[5:7] // reference to subarray {5,6} - len(a) is 2 and cap(a) is 5
	t.Log(a)
	a = a[0:4] // ref of subarray {5,6,7,8} - len(a) is now 4 but cap(a) is still 5
	t.Log(a)
}
