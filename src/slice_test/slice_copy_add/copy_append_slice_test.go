package slice_copy_add

import (
	"fmt"
	"testing"
)

func TestCopyAppend(t *testing.T) {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	/*
		append() 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片；追加的元素必须和原切片的元素是同类型。
		如果 s 的容量不足以存储新增元素，append() 会分配新的切片来保证已有切片元素和新增元素的存储。
		返回的切片可能已经指向一个不同的相关数组了。append() 方法总是返回成功，除非系统内存耗尽了

	*/
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

/*
该方法位append的相关类似源代码
func copy(dst, src []T) int 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst，覆盖 dst 的相关元素，并且返回拷贝的元素个数。
源地址和目标地址可能会有重叠。拷贝个数是 src 和 dst 的长度最小值。如果 src 是字符串那么元素类型就是 byte。如果继续使用 src，在拷贝结束后执行 src = dst。
*/
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]

	copy(slice[m:n], data)
	return slice
}
