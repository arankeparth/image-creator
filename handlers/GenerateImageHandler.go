package handlers

import (
	"encoding/json"
	"image-creator/spec"
	"image-creator/codeToImage"
	"log"
	"net/http"
)

const (
	url = "https://login.cloudtvos.com/"
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
	content, err := codeToImage.GenerateImage(req.IsUpdate, req.Code, url)
	if err != nil {
		log.Printf("Failed to generate image: %s\n", err.Error())
		http.Error(w, "Failed to generate image", http.StatusBadRequest)
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
