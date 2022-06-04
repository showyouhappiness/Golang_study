package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

func QrcodesController(c *gin.Context) { // 定义一个函数，用于处理二维码请求
	if content := c.Query("content"); content != "" { // 获取请求中的参数,并判断是否为空，如果不为空，则进行二维码生成
		png, err := qrcode.Encode(content, qrcode.Medium, 256) // 生成二维码
		if err != nil {
			log.Fatal(err)
		}
		c.Data(http.StatusOK, "image/png", png) // 返回二维码
	} else { // 如果为空，则返回错误码
		c.Status(http.StatusPreconditionRequired) // 返回错误码
	}
}
