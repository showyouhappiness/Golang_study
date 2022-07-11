/*
基于结构体实现了一个栈结构，为栈的实现创建一个单独的包 stack，并从 main 包 main.stack.go 中调用它。
*/
package method

import (
	"fmt"
	"strconv"
	"testing"
)

const LIMIT = 4

type Stack struct {
	ix   int // first free position, so data[ix] == 0
	data [LIMIT]int
}

type Stack1 [LIMIT]int

func TestStackStruct(t *testing.T) {
	st1 := new(Stack)
	fmt.Printf("%v\n", st1)
	st1.Push(3)
	fmt.Printf("%v\n", st1)
	st1.Push(7)
	fmt.Printf("%v\n", st1)
	st1.Push(10)
	fmt.Printf("%v\n", st1)
	st1.Push(99)
	fmt.Printf("%v\n", st1)
	p := st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop()
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
}

func (st *Stack) Push(n int) {
	if st.ix+1 > LIMIT {
		return // stack is full!
	}
	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) Pop() int {
	st.ix--
	return st.data[st.ix]
}

func (st1 *Stack1) Pop1() int {
	v := 0
	for ix := len(st1) - 1; ix >= 0; ix-- {
		if v = st1[ix]; v != 0 {
			st1[ix] = 0
			return v
		}
	}
	return 0
}

func (st Stack) String() string {
	str := ""
	for ix := 0; ix < st.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "] "
	}
	return str
}
