// 匿名函数同样被称之为闭包（函数式语言的术语）：它们被允许调用定义在其它环境下的变量。
//闭包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态。另一种表示方式为：一个闭包继承了函数所声明时的作用域。
//这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁。
//闭包经常被用作包装函数：它们会预先定义好 1 个或多个参数以用于包装，另一个不错的应用就是使用闭包来完成更加简洁的错误检查。
package closure_test

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"testing"
	"time"
)

func work() func() int {
	days := 0
	return func() int {
		days++
		return days
	}
}

func TestClosure(t *testing.T) {
	closure := work()
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
}

func func1() (i int) {
	i = 100
	defer func() {
		i += 1
	}()
	return 5
}

func func2() int {
	var i int
	defer func() {
		i += 1

	}()
	i += 100

	return i
}

func TestClosure1(t *testing.T) {
	fmt.Println(func1())
	fmt.Println(func2())
}

func TestClosure2(t *testing.T) {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func TestClosure3(t *testing.T) {
	var f = Adder1()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder1() func(int) int {
	var x int
	return func(delta int) int {
		println(delta)
		x += delta
		return x
	}
}

func TestClosure4(t *testing.T) {
	fmt.Println(f())
}
func f() (ret int) {
	// 关键字 defer经常配合匿名函数使用，它可以用于改变函数的命名返回值。
	// 匿名函数还可以配合 go 关键字来作为 goroutine 使用
	defer func() {
		ret++
	}()
	return 1
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func TestClosure5(t *testing.T) {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	println(addBmp("file"))
	println(addJpeg("file"))
}

// 闭包调试，可以获取哪个文件中的具体哪个函数正在执行

// 包 runtime 中的函数 Caller() 提供了相应的信息，因此可以在需要的时候实现一个 where() 闭包函数来打印函数执行的位置
func TestRunTime(t *testing.T) {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	// some code
	where()
	// some more code
	where()
}

// 设置 log 包中的 flag 参数来实现
func TestSetFlags(t *testing.T) {
	log.SetFlags(log.Llongfile)
	log.Print("")

	//简短版本的 where() 函数
	var where = log.Print
	func() {
		where()
		// some code
		where()
		// some code
		where()
	}()
}

// 计算函数执行时间
func TestGetRunTime(t *testing.T) {
	start := time.Now()
	time.Sleep(time.Second * 2)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
