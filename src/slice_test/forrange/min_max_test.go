/*写一个 minSlice() 方法，传入一个 int 的切片并且返回最小值，再写一个 maxSlice() 方法返回最大值。*/
package forrange

import (
	"fmt"
	"math"
	"testing"
)

func TestMinMax(t *testing.T) {
	sl1 := []int{78, 34, 643, 12, 90, 492, 13, 2}
	max := maxSlice(sl1)
	fmt.Printf("The maximum is %d\n", max)
	min := minSlice(sl1)
	fmt.Printf("The minimum is %d\n", min)
}

func maxSlice(sl []int) (max int) {
	for _, v := range sl {
		if v > max {
			max = v
		}
	}
	return
}

func minSlice(sl []int) (min int) {
	// min = int(^uint(0) >> 1)
	min = math.MaxInt32
	for _, v := range sl {
		if v < min {
			min = v
		}
	}
	return
}

/* Output:
The maximum is 643
The minimum is 2
*/
