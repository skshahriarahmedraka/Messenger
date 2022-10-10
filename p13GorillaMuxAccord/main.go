package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"app/config"
	// "app/logerror"
	"app/route"

	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
	"github.com/pytimer/mux-logrus"
)

// func ERROR(e error, s string ){
// 	if e!= nil {
// 		log.Fatalln(s)
// 	}
// }

func init(){

	// LOAD ENVIRONMENT VARIABLES
	config.LoadEnvironmentVar()

}
func main() {
	//LOAD ENVIRONMENT VARIABLES 
	// err:=godotenv.Load()
    // logerror.ERROR("🔥❌ ~ file: main.go ~ line 24 ~ funcmain ~ err : ",err)
	

	// CREATE GORILLA MUX WRITER 
    r := mux.NewRouter()
	route.Router(r)
    srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8888",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	r.Use(muxlogrus.NewLogger().Middleware)
	fmt.Println("✨ Server listening on localhost:8888 ... ")
	if err:= srv.ListenAndServe();err!=nil {
		log.Fatalln(err)
	}
}