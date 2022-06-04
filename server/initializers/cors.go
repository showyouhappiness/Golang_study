package initializers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func InitCors(router *gin.Engine) { // 初始化跨域
	router.Use(cors.New(cors.Config{ // 创建一个跨域配置
		AllowMethods:     []string{"PUT", "PATCH", "POST"},   // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Content-Type"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},         // 暴露的请求头
		AllowCredentials: true,                               // 是否允许携带cookie
		AllowOriginFunc: func(origin string) bool { // 允许的域名
			if origin == "http://127.0.0.1:27149" || origin == "http://:localhost:27149" {
				return true
			} else {
				log.Printf("%v is now allowed", origin) // 打印被拒绝的域名
				return false
			}
		},
	}))
}
