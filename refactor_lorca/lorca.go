package refactor_lorca

import (
	"github.com/gin-gonic/gin"
	"os/exec"
	"time"
)

func Lorca2() {
	go func() { // 开启一个goroutine,如果不将gin.Default()放在goroutine中，那么gin.Default()会在主goroutine中执行，会导致主线程阻塞
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.String(200, "Hello World")
		})
		r.Run(":8080")
	}()

	chromePath := "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://localhost:8080")
	cmd.Start()
	select {
	case <-time.After(time.Second * 5):
		cmd.Process.Kill()
	}
}
