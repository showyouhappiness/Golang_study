package config

import (
	"log"
	"os"
	"path/filepath"
)

func DefaultPathes() (uploads string) {
	exe, err := os.Executable() // 获取当前执行文件的路径
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)                // 获取当前执行文件的父目录
	uploads = filepath.Join(dir, "uploads") // 拼接路径
	return
}

var UploadsDir = DefaultPathes()
