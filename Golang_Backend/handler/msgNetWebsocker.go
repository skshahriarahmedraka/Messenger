package handler

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (H *DatabaseCollections) MsgNetSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
    }
	c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer c.Close()
	
}
