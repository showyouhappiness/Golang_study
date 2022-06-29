/*
遍历一个字符数组，并将当前字符和前一个字符不相同的字符拷贝至另一个数组。
*/
package string

import (
	"fmt"
	"testing"
)

var arr []byte = []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd', 'e', 'f', 'g'}

func TestUniq(t *testing.T) {
	arru := make([]byte, len(arr)) // this will contain the unique items
	ixu := 0                       // index in arru
	tmp := byte(0)
	for _, val := range arr {
		if val != tmp {
			arru[ixu] = val
			fmt.Printf("%c ", arru[ixu])
			ixu++
		}
		tmp = val
	}
	// fmt.Println(arru)
}
