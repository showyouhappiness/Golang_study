package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func FilesController(c *gin.Context) {
	file, err := c.FormFile("raw") //获取上传的文件
	if err != nil {
		log.Fatal(err)
	}
	exe, err := os.Executable() //获取当前执行文件的路径
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)                 //获取当前执行文件的父目录
	uploads := filepath.Join(dir, "uploads") //拼接路径
	filename := uuid.New().String()          //生成文件名
	err = os.MkdirAll(uploads, os.ModePerm)  //创建目录
	if err != nil {
		log.Fatal(err)
	}
	fullpath := path.Join("uploads", filename+filepath.Ext(file.Filename)) //拼接路径
	fileErr := c.SaveUploadedFile(file, filepath.Join(dir, fullpath))      //保存文件
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath}) //返回文件路径
}
