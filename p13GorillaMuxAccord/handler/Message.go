package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)
func (H *DatabaseCollections) Message(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// hj, ok := w.(http.Hijacker)
	// if !ok {
	// 	http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
	// 	return
	// }
	// conn, bufrw, err := hj.Hijack()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// // Don't forget to close the connection:
	// defer conn.Close()
	// bufrw.WriteString("Now we're speaking raw TCP. Say hi: ")
	// bufrw.Flush()
	// s, err := bufrw.ReadString('\n')
	// if err != nil {
	// 	log.Printf("error reading string: %v", err)
	// 	return
	// }
	// fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
	// bufrw.Flush()

	fmt.Printf("w's type is %T\n", w)


	vars := mux.Vars(r)
	SendID := vars["SendID"]
	fmt.Println("🚀 ~ file: Message.go ~ line 14 ~ func ~ SendID : ", SendID)
	ReceiverID := vars["ReceiverID"]
	fmt.Println("🚀 ~ file: Message.go ~ line 16 ~ func ~ ReceiverID : ", ReceiverID)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	// var origins = []string{ "http://127.0.0.1:3000", "http://localhost:18081", "http://example.com"}
	

	conn2, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("❌ ~ file: Message.go ~ line 17 ~ func ~ err : ", err)
		
	}
	defer conn2.Close()

	for {
		messageType, p, err := conn2.ReadMessage()
		fmt.Println("🚀 ~ file: Message.go ~ line 39 ~ func ~ p : ", string(p))
		if err != nil {
			fmt.Println("❌ ~ file: Message.go ~ line 26 ~ func ~ err : ", err)
			return
		}
		if err := conn2.WriteMessage(messageType, p); err != nil {
			fmt.Println("❌ ~ file: Message.go ~ line 31 ~ iferr:=conn.WriteMessage ~ err : ", err)
			return
		}
	}

}
