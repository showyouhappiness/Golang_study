package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func TextsController(c *gin.Context) {
	var json struct { //定义结构体
		Raw string //定义字段
	}
	if err := c.ShouldBindJSON(&json); err != nil { //获取json数据
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //返回错误
	} else {
		exe, err := os.Executable() //获取当前执行文件的路径
		if err != nil {
			log.Fatal(err)
		}
		dir := filepath.Dir(exe) //获取当前执行文件的父目录
		if err != nil {
			log.Fatal(err)
		}
		filename := uuid.New().String()          //生成文件名
		uploads := filepath.Join(dir, "uploads") //拼接路径
		err = os.MkdirAll(uploads, os.ModePerm)  //创建目录
		if err != nil {
			log.Fatal(err)
		}
		filePath := path.Join("uploads", filename+".txt")                            //拼接路径
		err = ioutil.WriteFile(filepath.Join(dir, filePath), []byte(json.Raw), 0644) //写入文件
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"url": "/" + filePath}) //返回json数据
	}
}
