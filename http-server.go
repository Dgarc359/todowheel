package main

import (
	"fmt"
	_ "io"
	"log"
	"net/http"
	util "todowheel-backend/http-methods"
)

/*
Starts HTTP server on specified port
*/
func startServer() {

    http.HandleFunc("/health", util.GetHealth)
    
    //http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    //    util.GetHealth(w, r)
    //})

    fmt.Printf("Server running (port=8080), route: http://localhost:8080/health\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
