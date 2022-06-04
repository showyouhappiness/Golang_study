package controllers

import (
	"Golang_study/server/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func UploadsController(c *gin.Context) {
	if path := c.Param("path"); path != "" { // 获取路径，如果路径不为空
		target := filepath.Join(config.UploadsDir, path)             // 拼接路径
		c.Header("Content-Description", "File Transfer")             // 设置文件头
		c.Header("Content-Transfer-Encoding", "binary")              // 设置文件编码
		c.Header("Content_Disposition", "attachment;filename="+path) // 设置文件名
		c.Header("Content-Type", "application/octet-stream")         // 设置文件类型
		c.File(target)                                               // 返回文件
	} else {
		c.Status(http.StatusNotFound) // 如果路径为空，返回404
	}
}
