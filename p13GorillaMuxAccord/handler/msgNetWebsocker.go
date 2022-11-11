package handler

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func (H *DatabaseCollections) MsgNetSocket(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		fmt.Println("🚀 ~ file: Message.go ~ line 39 ~ func ~ p : ", string(p))
		if err != nil {
			fmt.Println("❌ ~ file: Message.go ~ line 26 ~ func ~ err : ", err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println("❌ ~ file: Message.go ~ line 31 ~ iferr:=conn.WriteMessage ~ err : ", err)
			return
		}
	}
}
