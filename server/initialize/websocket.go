package initialize

import (
	"github.com/gorilla/websocket"
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// 允许所有请求来源
			return true
		},
	}

	clients = make(map[*websocket.Conn]bool)
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	// 将新的连接加入客户端列表
	clients[conn] = true

	for {
		// 读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			delete(clients, conn)
			break
		}

		global.LOG.Println(string(msg))

		// 将消息广播给所有连接的客户端
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Failed to send message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
