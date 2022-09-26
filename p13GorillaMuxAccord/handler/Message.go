package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func(H *DatabaseCollections) Message(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	SendID := vars["SendID"]
    fmt.Println("🚀 ~ file: Message.go ~ line 14 ~ func ~ SendID : ", SendID)
	ReceiverID := vars["ReceiverID"]
    fmt.Println("🚀 ~ file: Message.go ~ line 16 ~ func ~ ReceiverID : ", ReceiverID)

	

	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Add("Content-Type","application/json")

	// var origins = []string{ "http://127.0.0.1:3000", "http://localhost:18081", "http://example.com"}
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn,err:= upgrader.Upgrade(w,r,nil)
	if err!=nil {
		fmt.Println("❌ ~ file: Message.go ~ line 17 ~ func ~ err : ", err)
		
	}

	for {
		messageType,p,err:= conn.ReadMessage()
        fmt.Println("🚀 ~ file: Message.go ~ line 39 ~ func ~ p : ", string(p))
		if err !=nil {
			fmt.Println("❌ ~ file: Message.go ~ line 26 ~ func ~ err : ", err)
			return 
		}
		if err := conn.WriteMessage(messageType,p);err !=nil {
            fmt.Println("❌ ~ file: Message.go ~ line 31 ~ iferr:=conn.WriteMessage ~ err : ", err)
			return 
		}
	}

}


