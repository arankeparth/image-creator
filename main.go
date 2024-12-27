package main

import (
	"image-creator/handlers"
	"image-creator/spec"
	"log"
	"net/http"
)

func main() {
	// Register HTTP handlers
	http.HandleFunc(spec.GenerateImagePath, handlers.GenerateImageHandler)
	// Start the HTTP server
	port := ":8081"
	log.Printf("Server is running on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
