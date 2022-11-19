package router

import (
	// "app/controller"
	// "app/database"
	// "fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func RouteWithoutAuth(r *gin.Engine) {
	// fmt.Println("🚀 ~ file: routers.go ~ line 21 ~ funcUserRoutes ~ database.Client : ", database.Client)
	// DB2 := database.MongodbConnection()
	// H := database.DatabaseInitialization(DB2)
	// fmt.Println("🚀 ~ file: WithoutAuth.go ~ line 13 ~ funcRouteWithoutAuth ~ RouteWithoutAuth  ", H)
	r.GET("/gin/ws", Instantbuy())
	// r.POST("/sveltekit/register", H.SveltekitRegister())
	// r.POST("/user/login", H.Login())
	// r.POST("/user/register", H.Register())
	// r.POST("/user/profileupdate", H.ProfileUpdate())

}
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Instantbuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
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
}