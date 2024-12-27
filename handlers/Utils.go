package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image-creator/codeToImage/codeToImage"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func getImage(isUpdate bool) (string, error) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random 4-digit code
	code := rand.Intn(9000) + 1000 // Ensures the number is between 1000 and 9999

	// Generate an image using the codeToImage package
	imagePath, err := codeToImage.CodeToImage(isUpdate, fmt.Sprintf("%d", code), url)
	if err != nil {
		fmt.Printf("Error generating image: %v\n", err)
		return "", err
	}
	// Open the generated image file
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	// Read the image file into a byte slice (all binary data)
	imageData, err := ioutil.ReadAll(imageFile)
	if err != nil {
		fmt.Printf("Error reading image file: %v\n", err)
		return "", err

	}

	// Encode the binary data of the image into Base64 (token format)
	token := base64.StdEncoding.EncodeToString(imageData)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}

	fmt.Printf("Successfully generated token for code: %d\n", code)

	return token, nil
}

func writeResponse(w http.ResponseWriter, isUpdate bool) {
	// Create a map or struct to hold the JSON response
	content, err := getImage(isUpdate)

	if err != nil {
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
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
