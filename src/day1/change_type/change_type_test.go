package change_type

import "testing"

//定义一个别名
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	//b = a //如果不显式转化就会报错，因为Go语言不支持隐式类型转换
	b = int64(a)
	t.Log(a, b)
	var c MyInt
	//c = b //别名做隐式类型转换也是不行的
	c = MyInt(b)
	t.Log(c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1 // Go不支持指针运算，所以这块会报错，一般我们会进行指针的运算（指针的自针），访问连续的存储空间（例如数组）
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string //声明的会是一个空值 nil,实际上string在Golang里面是一个数值类型，所以它初始化零值的时候，会初始化空字符串
	t.Log("000" + s + "000")
	t.Log(len(s))
	//if s == nil {
	//}
	//if s == "" {
	//}
}
