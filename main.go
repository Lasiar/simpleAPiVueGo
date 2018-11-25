package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// create custom serveMux
	apiMux := http.NewServeMux()

	// register handle
	apiMux.HandleFunc("/api/get-server-time", HandleGetServerTime)

	fmt.Println("server listen on ", GetConfig().Port)

	err := http.ListenAndServe(":"+GetConfig().Port, MiddlewareCORS(apiMux))
	if err != nil {
		log.Fatal(err)
	}
}
