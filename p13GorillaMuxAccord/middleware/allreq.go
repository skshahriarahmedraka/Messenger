package middleware

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
