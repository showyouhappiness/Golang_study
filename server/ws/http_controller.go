package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var wsupgrader = websocket.Upgrader{ // 定义升级协议的参数
	ReadBufferSize:  1024, // 设置输入缓冲区大小
	WriteBufferSize: 1024, // 设置输出缓冲区大小
	CheckOrigin: func(r *http.Request) bool { // 设置检查来源的函数
		return true
	},
}

func wshandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	connect, err := wsupgrader.Upgrade(w, r, nil) // 完成ws升级
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: connect, send: make(chan []byte, 256)} // 定义客户端
	client.hub.register <- client                                            // 注册到hub
	go client.writePump()                                                    // 开启发送协程
	go client.readPump()                                                     // 开启接收协程
}

func HttpController(c *gin.Context, hub *Hub) {
	wshandler(hub, c.Writer, c.Request)
}
