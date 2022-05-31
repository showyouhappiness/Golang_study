package gin_study

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

// go:embed frontend/dist/*
var FS embed.FS

func Gin1() {
	go func() { // 开启一个goroutine,如果不将gin.Default()放在goroutine中，那么gin.Default()会在主goroutine中执行，会导致主线程阻塞
		gin.SetMode(gin.DebugMode)
		r := gin.Default()
		staticFiles, _ := fs.Sub(FS, "frontend/dist")
		r.StaticFS("/static", http.FS(staticFiles))
		r.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path
			if strings.HasPrefix(path, "/static/") {
				reader, err := staticFiles.Open("index.html")
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()       // 关闭文件
				stat, err := reader.Stat() // 获取文件信息
				if err != nil {
					log.Fatal(err)
				}
				c.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
			} else {
				c.Status(http.StatusNotFound)
			}
		})
		r.Run(":8080")
	}()
	chSignal := make(chan os.Signal, 1)   //创建一个信号通道
	signal.Notify(chSignal, os.Interrupt) //监听指定信号
	fmt.Println("server start")
	chromePath := "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080/static/index.html")
	cmd.Start()

	select {
	case <-chSignal: //阻塞直到接收到信号
		cmd.Process.Kill()
	}

}
