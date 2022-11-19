package middleware

import (
	"app/logerror"
	"app/model"
	// "encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// import (
// 	"os"
// 	"time"

// 	"net/http"

// 	logg "github.com/sirupsen/logrus"
// )
// func init() {
// 	logg.SetOutput(os.Stdout)
// }
// func ReqLog(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		logg.WithFields(logg.Fields{
// 			"method":     r.Method,
// 			"path":       r.URL.Path,
// 			"status":     r.Response.StatusCode,
// 			"latency_ns": time.Since(start).Nanoseconds(),
// 		}).Info("request details")
// 		 h.ServeHTTP(w, r)
// 	})
// }


func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter , r *http.Request){
		
		w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Add("Content-Type","application/json")
	// an example API handler
	fmt.Println("Auth middleWare has called ...")
	cookie, err := r.Cookie("Auth1")

	if err != nil {
		// cookie2, err2 := r.Cookie("RefAuth1")
		// if err2 !=nil {

		// }
		if err == http.ErrNoCookie {
			logerror.ERROR("🚀 ~ file: auth.go ~ line 49 ~ returnhttp.HandlerFunc ~ err : ", err)
			// w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w,r,"/register",http.StatusSeeOther)
			return
		}
		logerror.ERROR("🚀 ~ file: auth.go ~ line 49 ~ returnhttp.HandlerFunc ~ err : ", err)
		// w.WriteHeader(http.StatusBadRequest)
		http.Redirect(w,r,"/register",http.StatusSeeOther)
		return 
	}

	tokenStr:= cookie.Value
    fmt.Println("🚀 ~ file: home.go ~ line 28 ~ func ~ tokenStr : ", tokenStr)

	claims := model.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		fmt.Println("🚀 ~ file: home.go ~ line 34 ~ func ~ err : ", err)
	if err != nil {
		if err!= jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	fmt.Println("🚀 ~ file: home.go ~ line 60 ~ func ~ tkn.Valid : ", tkn.Valid)
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
	}
	fmt.Println("✨Token is valid Welcome home")
	next.ServeHTTP(w,r)

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(`{status: "Welcome home"}`)
		
	})
}

func RefCookieValidityCheck(){

}