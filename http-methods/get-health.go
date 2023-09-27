package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
	body, err := io.ReadAll(r.Body)
	content := r.Header.Get("Content-Type")
	fmt.Printf("\nContent %s\n", content)

	if err != nil {
		log.Printf("Error reading body, %v", err)
		return
	}

	fmt.Printf("Request body, %s", body)

}
