package main

import (
	"fmt"
	"sync"
	"time"
)

type Glimit struct {
	n int
	c chan struct{}
}

// initialization Glimit struct
func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

// Run f in a new goroutine but with limit.
func (g *Glimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}

var wg = sync.WaitGroup{}

// 有限制
func main() {
	number := 10000
	g := New(10)
	for i := 0; i < number; i++ {
		wg.Add(1)
		value := i
		goFunc := func() {
			// 做一些业务逻辑处理
			fmt.Printf("go func: %d\n", value)
			time.Sleep(time.Second)
			wg.Done()
		}
		g.Run(goFunc)
	}
	wg.Wait()
}

// 无限制
func main1() {
	number := 10000
	for i := 0; i < number; i++ {
		value := i
		go func() {
			// 做一些业务逻辑处理
			fmt.Printf("go func: %d\n", value)
		}()
	}
}
