package server

import (
	"Golang_study/server/controllers"
	"Golang_study/server/initializers"
	"Golang_study/server/ws"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"strings"
)

// Fs 将所有文件打包到一个文件中
//go:embed server/frontend/dist/*
var Fs embed.FS

func Run(start chan int, end chan interface{}) {
	gin.SetMode(gin.ReleaseMode)  // 设置gin模式
	gin.DisableConsoleColor()     // 关闭gin控制台颜色
	router := gin.Default()       // 创建一个gin路由
	initializers.InitCors(router) // 初始化跨域
	hub := ws.NewHub()            // 创建一个websocket hub
	go hub.Run()                  // 启动websocket hub
	router.GET("/ws", func(c *gin.Context) {
		ws.HttpController(c, hub) // 启动http controller
	})
	router.GET("/uploads/:path", controllers.UploadsController)      // 启动上传文件的controller
	router.GET("/api/v1/addresses", controllers.AddressesController) // 启动获取地址的controller
	router.GET("", controllers.QrcodesController)                    // 启动获取二维码的controller
	router.POST("/api/v1/files", controllers.FilesController)        // 启动上传文件的controller
	router.POST("/api/v1/texts", controllers.TextsController)        // 启动上传文本的controller
	staticFiles, _ := fs.Sub(Fs, "frontend/dist")                    // 获取静态文件
	router.StaticFS("/static", http.FS(staticFiles))                 // 启动静态文件
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path               // 获取请求路径
		if strings.HasPrefix(path, "/static/") { // 判断请求路径是否以/static/开头
			reader, err := staticFiles.Open("index.html") // 打开静态文件的index.html
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()       // 关闭reader
			stat, err := reader.Stat() // 获取文件的状态
			if err != nil {
				log.Fatal(err)
			}
			c.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", reader, nil) // 将文件内容写入response
		} else { // 如果不是以/static/开头
			c.Status(http.StatusNotFound) // 返回404
		}
	})
	port := 27149
	start <- port
	runErr := router.Run(fmt.Sprintf("%d", port)) // 启动服务'
	if runErr != nil {
		end <- runErr
		log.Fatal(runErr)
	}
}
