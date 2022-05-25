package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World v20!")
	})
	if os.Getenv("PORT") != "" {
		port := os.Getenv("PORT")
		fmt.Printf("Server running (port=%s), route: http://localhost:%s/\n", port, port)
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	} else {
		fmt.Printf("Server running (port=8080), route: http://localhost:8080/\n")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
