/*
编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
*/
package string

import (
	"fmt"
	"testing"
)

func TestSplitString(t *testing.T) {
	str := "Google"
	for i := 0; i <= len(str); i++ {
		a, b := Split(str, i)
		fmt.Printf("The string %s split at position %d is: %s / %s\n", str, i, a, b)
	}

}

func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

/*
假设有字符串 str，返回 str[len(str)/2:] + str[:len(str)/2] 的结果
*/

func TestSplitString2(t *testing.T) {
	str := "Google"
	str2 := Split2(str)
	fmt.Printf("The string %s transformed is: %s\n", str, str2)
}

func Split2(s string) string {
	mid := len(s) / 2
	return s[mid:] + s[:mid]
}
