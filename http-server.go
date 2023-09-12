package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
Starts HTTP server on specified port
*/
func startServer() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "OK")
        body, err := io.ReadAll(r.Body)
        content := r.Header.Get("Content-Type")
        fmt.Printf("\nContent %s\n", content)

        if err != nil {
            log.Printf("Error reading body, %v", err)
            return
        }
        fmt.Printf("Request body, %s", body)
    })

    fmt.Printf("Server running (port=8080), route: http://localhost:8080/health\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
