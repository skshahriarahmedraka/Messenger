package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"app/config"
	// "app/logerror"
	"app/route"

	// "github.com/gin-contrib/cors"
	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
	"github.com/pytimer/mux-logrus"
	"github.com/rs/cors"
	// "app/cmd"
)

// func ERROR(e error, s string ){
// 	if e!= nil {
// 		log.Fatalln(s)
// 	}
// }

func init() {

	// LOAD ENVIRONMENT VARIABLES
	config.LoadEnvironmentVar()

}
func main() {

	// go cmd.Main()
	//LOAD ENVIRONMENT VARIABLES
	// err:=godotenv.Load()
	// logerror.ERROR("🔥❌ ~ file: main.go ~ line 24 ~ funcmain ~ err : ",err)

	// CREATE GORILLA MUX WRITER
	r := mux.NewRouter()
	route.Router(r)

	// handler := cors.Default().Handler(r)
	// c:=cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedCredentials: true,
	// })
	// handler = c.Handler(handler)

	handler := cors.Default().Handler(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	// decorate existing handler with cors functionality set in c
	handler = c.Handler(handler)

	// log.Println("Serving at localhost:5000...")
	// log.Fatal(http.ListenAndServe(":5000", handler))

	srv := &http.Server{
		Handler: handler,
		Addr:    "127.0.0.1:8888",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	r.Use(muxlogrus.NewLogger().Middleware)

	fmt.Println("✨ Server listening on localhost:8888 ... ")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
