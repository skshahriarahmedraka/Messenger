package handler

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func (H *DatabaseCollections) MsgGobwas(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		fmt.Println("🚀 ~ file: msgGobwas.go ~ line 13 ~ func ~ err : ", err)
		// handle error
	}
	go func() {
		defer conn.Close()

		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				fmt.Println("🚀 ~ file: msgGobwas.go ~ line 22 ~ gofunc ~ err : ", err)
				// handle error
			}
			err = wsutil.WriteServerMessage(conn, op, msg)
			if err != nil {
				fmt.Println("🚀 ~ file: msgGobwas.go ~ line 27 ~ gofunc ~ err : ", err)
				// handle error
			}
		}
	}()
}
