/*
使用冒泡排序的方法排序一个包含整数的切片
*/
package string

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	sla := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	fmt.Println("before sort: ", sla)
	// sla is passed via call by value, but since sla is a reference type
	// the underlying slice is array is changed (sorted)
	bubbleSort(sla)
	fmt.Println("after sort:  ", sla)
}

func bubbleSort(sl []int) {
	// passes through the slice:
	for pass := 1; pass < len(sl); pass++ {
		// one pass:
		for i := 0; i < len(sl)-pass; i++ { // the bigger value 'bubbles up' to the last position
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
			}
		}
	}
}
