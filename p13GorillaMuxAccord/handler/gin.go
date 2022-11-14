package handler
import (
	"fmt"
	"log"
	"os"

	// "app/config"
	// "app/controller"
	// "app/middleware"
	// "app/cmd/router"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
)

func Main(H *DatabaseCollections) {
	gin.SetMode(gin.ReleaseMode)

	//  fmt.Println("environment var :  ", os.Getenv("NAME"))
	//  fmt.Println("environment var :  ", os.Getenv("NAME")=="SK SHAHRIAR AHMED RAKA")
	// fmt.Println("environment var :", os.Getenv("POSTGRES_TIMEZONE"))
	// fmt.Println("environment var :", os.Getenv("POSTGRES_TIMEZONE")=="Asia/Dhaka")
	fmt.Println("🚀✨ Api is started")
	// r := gin.Default()

	r := gin.New()
	
	// r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())

	// router.RouteWithoutAuth(r)
	r.GET("/gin/hello", H.Hello())
	r.GET("/gin/ws", H.Instantbuy())

	log.Println("Server is started in PORT 8889 ...👨‍💻 ")
	if e := r.Run(os.Getenv("HOST") + ":" + os.Getenv("GIN_PORT")); e != nil {
		log.Fatalln("❌ ERROR when Server is start   👨‍💻 : ", e)
	}
}