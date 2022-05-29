package gin_study

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
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
		err := r.Run(":8080")
		if err != nil {
			return
		}
	}()
	chSignal := make(chan os.Signal, 1)                      //创建一个信号通道
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM) //监听指定信号
	fmt.Println("server start")
	chromePath := "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://localhost:8080")
	cmd.Start()
	select {
	case <-time.After(time.Second * 5):
		cmd.Process.Kill()
	}

}
