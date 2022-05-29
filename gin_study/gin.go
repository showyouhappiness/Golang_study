package gin_study

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"os/signal"
)

func Gin1() {
	go func() { // 开启一个goroutine,如果不将gin.Default()放在goroutine中，那么gin.Default()会在主goroutine中执行，会导致主线程阻塞
		gin.SetMode(gin.DebugMode)
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run(":8080")

	}()
	chSignal := make(chan os.Signal, 1)   //创建一个信号通道
	signal.Notify(chSignal, os.Interrupt) //监听指定信号
	fmt.Println("server start")
	chromePath := "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://localhost:8080")
	cmd.Start()

	select {
	case <-chSignal: //阻塞直到接收到信号
		cmd.Process.Kill()
	}

}
