package handlers

import (
	"encoding/base64"
	"image-creator/codeToImage/codeToImage"
	"io/ioutil"
	"log"
	"os"
)

func getImage(isUpdate bool, code string) (string, error) {
	imagePath, err := codeToImage.GenerateImage(isUpdate, code, url)
	if err != nil {
		log.Printf("Error generating image: %v\n", err)
		return "", err
	}

	// Open the generated image file
	imageFile, err := os.Open(imagePath)
	if err != nil {
		log.Printf("Error opening image file: %v\n", err)
		return "", err
	}
	defer imageFile.Close()

	// Read the image file into a byte slice (all binary data)
	imageData, err := ioutil.ReadAll(imageFile)
	if err != nil {
		log.Printf("Error reading image file: %v\n", err)
		return "", err
	}

	// Encode the binary data of the image into Base64 (token format)
	token := base64.StdEncoding.EncodeToString(imageData)
	if err != nil && err.Error() != "EOF" {
		log.Printf("Error encoding image to Base64: %v\n", err)
		return "", err
	}

	// Delete the image file after generating the token
	err = os.Remove(imagePath)
	if err != nil {
		log.Printf("Error deleting image file: %v\n", err)
		return "", err
	}

	log.Printf("Successfully generated token for code: %s\n", code)

	return token, nil
}
