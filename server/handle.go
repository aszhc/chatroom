package server

import (
	"net/http"

	"github.com/aszhc/chatroom/logic"
)

var rootDir string

func RegisterHandle() {
	// 广播消息处理
	go logic.Broadcaster.Start() // 启动一个goroutine广播

	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/user_list", userListHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)
}
