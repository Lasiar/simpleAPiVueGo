package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// MiddlewareCORS set cors headers and cancel options request
func MiddlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set header CORS
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		// to cancel preflight xss
		if r.Method == http.MethodOptions {
			return
		}

		// handle next
		next.ServeHTTP(w, r)
	})
}

//  HandleGetServerTime simple work
func HandleGetServerTime(w http.ResponseWriter, _ *http.Request) {

	// local struct for response
	response := struct {
		Time     string `json:"time"`
		HostName string `json:"hostname"`
	}{}

	// get local time
	response.Time = time.Now().Format("2006-01-02 15:04:05")

	// get hostName
	response.HostName, _ = os.Hostname()

	// create encoder
	encoder := json.NewEncoder(w)

	// send request
	if err := encoder.Encode(response); err != nil {
		log.Println(err)
	}
}
