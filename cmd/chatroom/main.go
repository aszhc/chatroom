package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aszhc/chatroom/global"

	"github.com/aszhc/chatroom/server"
)

var (
	addr   = ":2022"
	banner = "            ____ _           _\n" +
		"	   / ___| |__   __ _| |_\n" +
		"	  | |   | '_ \\ / _` | __|\n" +
		"	  | |___| | | | (_| | |_\n" +
		"	   \\____|_| |_|\\__,_|\\__|\n" +
		"ChatRoom Powered By WebSocket, start on: %s"
)

func init() {
	global.Init()
}

func main() {
	fmt.Printf(banner+"\n", addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
