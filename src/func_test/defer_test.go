package func_test

import (
	"fmt"
	"io"
	"log"
	"testing"
)

func TestDeferFunc(t *testing.T) {
	function1()
	f()
	doDBOperations()
	func1("hello")
}

func function1() {
	defer function2()
	fmt.Printf("In function1 at the top\n")
	function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}

// 使用 defer 的语句同样可以接受参数
// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func f() {
	i := 0
	defer fmt.Println(i)
	i++
	defer print(i)
	return
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
