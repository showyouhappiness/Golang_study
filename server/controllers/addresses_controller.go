package controllers

import (
	"github.com/gin-gonic/gin"
	"net"
)

func AddressesController(c *gin.Context) {
	addrs, _ := net.InterfaceAddrs() // 获取本机所有ip地址
	var ipList []string              // 定义一个空的字符串切片,用于存储ip地址
	for _, address := range addrs {  // 遍历所有ip地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() { // 判断是否是ip地址,并且不是回环地址
			if ipNet.IP.To4() != nil { // 判断是否是ipv4地址
				ipList = append(ipList, ipNet.IP.String()) // 将ipv4地址追加到ipList中
			}
		}
	}
}
