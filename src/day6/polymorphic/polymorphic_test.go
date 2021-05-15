package polymorphic

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}
type GoProgrammer struct {
}

func (gp *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (jp *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func WriteFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphic(t *testing.T) {
	//goProg:=&GoProgrammer{} //和下面的一样 Programmer是一个interface，只能只一个指针类型的实例，所以接收的时候要加上取址符
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
}
