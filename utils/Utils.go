package utils

import (
	"image-creator/codeToImage"
	"log"
)


func GetImage(isUpdate bool, code string) (string, error) {
	token, err := codeToImage.GenerateImage(isUpdate, code, url)
	if err != nil {
		log.Printf("Error generating image: %v\n", err)
		return "", err
	}

	return token, nil
}
