/*
给定结构体类型 T，给 T 定义 String()，使得 fmt.Printf("%v\n", t) 输出：7 / -2.350000 / "abc\tdef"。
*/
package method

import (
	"fmt"
	"testing"
)

type T struct {
	a int
	b float32
	c string
}

func TestTypeString(t *testing.T) {
	t1 := T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t1)
	typeString()
}
func typeString() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
}
func (t *T) String() string {
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}
