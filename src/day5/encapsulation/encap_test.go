package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

//并没有对象复制产生，实例的数据是存放在同一块内存地址上的
//func (e *Employee) String() string {
//	fmt.Printf("Address is %X", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.ID, e.Name, e.Age)
//}

//实例的数据被复制了一份，整个结构被复制了；会有空间上更大复制的开销
func (e Employee) String() string {
	fmt.Printf("Address is %X", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.ID, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 21}
	e2 := new(Employee) //返回指针
	e2.ID = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.ID)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %X\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
