package main

import (
	"Golang_study/api/user/initialize"
)

func main() {
	//1. 初始化logger
	initialize.InitLogger()

	//2. 初始化配置文件
	initialize.InitConfig()
}
