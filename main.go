package main

import (
	"fmt"
	"time"
	"ws_server_active_connecctions/ws"
)

func main() {
	server := ws.StartServer(messageHandler)

	for {
		server.WriteMessage([]byte("Hello"))
		time.Sleep(time.Duration(1000))
	}
}

func messageHandler(message []byte) {
	fmt.Println(string(message))
}
