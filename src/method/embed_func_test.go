package method

import (
	"fmt"
	"testing"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

type Customer2 struct {
	Name string
	Log
}

func TestEmbedFunc1(t *testing.T) {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	c.Log().Add("2 - After me the world will be a better place!")
	//fmt.Println(c.log)
	fmt.Println(c.Log())

}

func TestEmbedFunc2(t *testing.T) {
	c := &Customer2{"Barak Obama", Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)

}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

//func (c *Customer2) Log2() *Log {
//	return c.Log2()
//}

func (c *Customer2) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log.String())
}

/*
内嵌的类型不需要指针，Customer 也不需要 Add 方法，它使用 Log 的 Add 方法，Customer 有自己的 String 方法，
并且在它里面调用了 Log 的 String 方法。如果内嵌类型嵌入了其他类型，也是可以的，那些类型的方法可以直接在外层类型中使用。
因此一个好的策略是创建一些小的、可复用的类型作为一个工具箱，用于组成域类型。
*/
