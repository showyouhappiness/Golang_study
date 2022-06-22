// 回调  函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调
// 将函数作为参数的最好的例子是函数 strings.IndexFunc()
package callback

import (
	"fmt"
	"strings"
	"testing"
)

func TestCallBack(t *testing.T) {
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

// strings_map.go

func TestStringsMap(t *testing.T) {
	// 过滤掉扩展ascii码的字符  https://tool.ip138.com/ascii_code/ 可查询ascii码
	asciiOnly := func(c rune) rune {
		println(c)
		if c > 127 {
			return ' '
		}
		return c
	}
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich"))

	// 映射func(rune)符文参数定义需要替换原始字符的字符
	modified := func(r rune) rune {
		if r == 'e' {
			return '@'
		}
		return r
	}

	input := "Hello, Welcome to GeeksforGeeks"
	fmt.Println(input)

	// using the function
	result := strings.Map(modified, input)
	fmt.Println(result)

	// 用于删除字符串中的单词之间的空格
	transformed := func(r rune) rune {

		if r == ' ' {
			return 0
		}
		return r
	}

	input1 := "GeeksforGeeks is a computer science portal."
	fmt.Println(input1)

	// using the function
	output := strings.Map(transformed, input1)
	fmt.Println(output)
}
