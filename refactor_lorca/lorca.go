package refactor_lorca

import (
	"github.com/gin-gonic/gin"
	"os/exec"
	"time"
)

func Lorca2() {
	go func() {
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
