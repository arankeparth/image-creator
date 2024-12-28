package main

import (
	"image-creator/handlers"
	"image-creator/spec"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// Register HTTP handlers
	http.HandleFunc(spec.GenerateImagePath, handlers.GenerateImageHandler)

	// Start the HTTP server with profiling
	port := ":8081"
	go func() {
		log.Println(http.ListenAndServe(":6060", nil)) // pprof server
	}()
	log.Printf("Server is running on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
