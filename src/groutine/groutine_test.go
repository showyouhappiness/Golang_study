package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		//如果我们需要一个函数在协程里面运行，在函数的前面加go结尾调用这个函数，把i传进去
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}

func TestGoroutine2(t *testing.T) {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func TestGoroutine3(t *testing.T) {
	ch1 := make(chan int)
	go pump(ch1) // pump hangs
	go suck(ch1)
}

func suck(ch1 chan int) {
	for true {
		fmt.Println(<-ch1) // prints only 0
	}
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func f1(in chan int) {
	fmt.Println(<-in)
}

func TestGoroutine4(t *testing.T) {
	out := make(chan int, 1)
	out <- 2
	go f1(out) // 可以使用带缓冲的通道或者将消费者提前至生产者的前方或者将生产者也写成协程的形式
	//select {}  // 可以让 main 程序无限阻塞
}
