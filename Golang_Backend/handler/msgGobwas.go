package handler

import (
	"fmt"
	"net/http"
	// "strings"
	// "time"

	"github.com/gorilla/websocket"
)

func (H *DatabaseCollections) MsgGobwas(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	
	ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ws.Close()
		for {
			//Read Message from client
			mt, message, err := ws.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			//If client message is ping will return pong
			if string(message) == "ping" {
				message = []byte("pong")
			}
			//Response message to client
			err = ws.WriteMessage(mt, message)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
}

