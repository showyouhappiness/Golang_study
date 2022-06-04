package server

import (
	"Golang_study/server/controllers"
	"Golang_study/server/initializers"
	"Golang_study/server/ws"
	"embed"
	"github.com/gin-gonic/gin"
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
	router.GET("/uploads/:path", controllers.UploadsController)
	router.GET("/api/v1/addresses", controllers.AddressesController)

}
