package handler

import "net/http"


func(H *DatabaseCollections) ServerChannelSettingsEdit(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Add("Content-Type","application/json")
}