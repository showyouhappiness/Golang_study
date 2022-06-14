//--file-version和--product-version标志可以采用特殊值：git-tag. 这将检索当前标签git describe --tags并将其添加到可执行文件的文件属性中。使用go generate命令生成可执行文件时，这个标志将被忽略。
//go:generate go-winres make --product-version=git-tag
package main

import (
	"Golang_study/server"
	"fmt"
	"github.com/zserge/lorca"
	"os"
	"os/signal"
	"sync"
)

func recoverFromError() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

func main() {
	var endWaiter sync.WaitGroup                     // 用于等待程序结束
	endWaiter.Add(1)                                 // 增加一个等待程序结束的 goroutine
	start := make(chan int)                          // 开始信号
	end := make(chan interface{})                    // 结束信号
	go server.Run(start, end)                        // 启动服务器
	go func(start chan int, quit chan interface{}) { // 启动浏览器
		port := <-start                                                                                                                        // 等待服务器启动
		defer recoverFromError()                                                                                                               // 发生错误时自动恢复
		ui, _ := lorca.New(fmt.Sprintf("http://127.0.0.1:%d/static/index.html", port), "", 800, 600, "--disable-sync", " --disable-translate") // 启动浏览器
		defer ui.Close()
		quit <- <-ui.Done() // 等待浏览器结束
	}(start, end)
	signalChannel := make(chan os.Signal, 1)   // 创建信号通道
	signal.Notify(signalChannel, os.Interrupt) // 注册Ctrl+c中断信号
	select {                                   // 等待信号
	case <-signalChannel: // 接收到中断信号
		endWaiter.Done() // 发送结束信号
	case <-end: // 接收到结束信号
		endWaiter.Done() // 发送结束信号
	}
	endWaiter.Wait() // 等待程序结束
	// test
}
