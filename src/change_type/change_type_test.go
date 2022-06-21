package change_type

import (
	"fmt"
	"testing"
)

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

func TestIota(t *testing.T) {
	const (
		a = 100  //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f        //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	//100 100 100 ha ha ha ha 7 8
	/*
	   从iota起到遇到别的赋值前，会在第一个行索引按顺序加一（即第一个出现iota为第八行，索引值就为7）
	*/
}

func TestByte(t *testing.T) {
	const (
		B float64 = 1 << (iota * 10)
		KB
		MB
		GB
	)
	fmt.Println(B, KB, MB, GB)
}
