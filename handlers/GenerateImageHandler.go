package handlers

import (
	"encoding/json"
	"image-creator/spec"
	"log"
	"net/http"
)

func GenerateImageHandler(w http.ResponseWriter, r *http.Request) {
	var req spec.GenerateImageRequest

	// Parse JSON request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Failed to parse JSON request body: %v\n", err)
		http.Error(w, "Failed to parse JSON request body", http.StatusBadRequest)
		return
	}

	// Validate code
	if len(req.Code) != 4 {
		log.Printf("Invalid code length: %s\n", req.Code)
		http.Error(w, "Code must be a 4-digit number", http.StatusBadRequest)
		return
	}

	// Create a map or struct to hold the JSON response
	content, err := getImage(req.IsUpdate, req.Code)
	if err != nil {
		log.Printf("Failed to generate image: %v\n", err)
		http.Error(w, "Failed to generate image", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"token": content,
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode JSON response: %v\n", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
