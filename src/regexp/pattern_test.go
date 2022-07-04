package regexp

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func TestPattern(t *testing.T) {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则
	//查找匹配的子串
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}
	//查找匹配的子串的位置
	if ok, _ := regexp.MatchString(pat, searchIn); ok {
		fmt.Println("Match Found2!")
	}

	/*
		必须先将正则模式通过 Compile() 方法返回一个 Regexp 对象。然后我们将掌握一些匹配，查找，替换相关的功能
		Compile() 函数也可能返回一个错误，一般使用时忽略对错误的判断，因为确信正则表达式是有效的。
		当用户输入或从数据中获取正则表达式的时候，就有必要去检验它的正确性。
		另外也可以使用 MustCompile() 方法，它可以像 Compile() 方法一样检验正则的有效性，但是当正则不合法时程序将 panic()
	*/
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	re, _ := regexp.Compile(pat) // Compile() 方法返回一个 Regexp 对象，或者返回一个错误。
	//将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#") // ReplaceAllString() 方法返回一个新的字符串，其中所有匹配的子串都用 replacement 替换。
	fmt.Println(str)
	//参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f) // ReplaceAllStringFunc() 方法返回一个新的字符串，其中所有匹配的子串都用函数返回的字符串替换。
	fmt.Println(str2)
}
