package main

import "fmt"

// 修改字符串的错误示例
//func main() {
//	x := "text"
//	x[0] = "T" // error: cannot assign to x[0]
//	fmt.Println(x)
//}

// 修改示例
func main() {
	x := "text"
	xBytes := []byte(x)
	fmt.Println(xBytes, xBytes[0])
	xBytes[0] = 'T' // 注意此时的 T 是 rune 类型
	fmt.Println(xBytes, xBytes[0])
	x = string(xBytes)
	fmt.Println(x) // Text
}
