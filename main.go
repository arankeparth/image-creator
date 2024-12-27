package main

import (
	"fmt"
	"image-creator/handlers"
	"image-creator/spec"
	"net/http"
)

func main() {
	// Register HTTP handlers
	http.HandleFunc(spec.LoginPath, handlers.LoginHandler)
	http.HandleFunc(spec.UpdatePath, handlers.UpdateHandler)
	// Start the HTTP server
	port := ":8081"
	fmt.Printf("Server is running on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
