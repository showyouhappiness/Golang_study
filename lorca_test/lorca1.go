package lorca_test

import (
	"github.com/zserge/lorca"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Lorca1() {
	// Create UI with basic HTML passed via data URI
	ui, _ := lorca.New("www.baidu.com", "", 480, 320)        // 打开百度
	chSignal := make(chan os.Signal, 1)                      //创建一个信号通道
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM) //监听指定信号
	select {                                                 //阻塞等待信号(当前goroutine)
	case <-ui.Done(): //ui关闭
		log.Println("UI is closed")
	case <-chSignal: //接收到信号,获取是哪个信号（终端、终止）
		log.Println("Received signal")
	}
	ui.Close() //关闭ui
}
