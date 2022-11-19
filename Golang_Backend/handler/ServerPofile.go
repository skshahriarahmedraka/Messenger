package handler

import (
	"encoding/json"
	"net/http"
)


func(H *DatabaseCollections) ServerProfile(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(`{status: "Server Profile successfull"}`)
}