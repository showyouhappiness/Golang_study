/*
写一个函数 InsertStringSlice() 将切片插入到另一个切片的指定位置。
*/
package slice_copy_add

import (
	"fmt"
	"testing"
)

func TestInsertSLice(t *testing.T) {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := InsertStringSlice(s, in, 0)
	fmt.Println(res) // [A B C M N O P Q R]
	res = InsertStringSlice(s, in, 3)
	fmt.Println(res) // [M N O A B C P Q R]
}

func InsertStringSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	fmt.Println(at)
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}
