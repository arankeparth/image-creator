# new IMAGES in WEBP 

## Generate Image

### For Create New User
```go
GenerateImage(isUpdate bool, cwToken string, valueofURL string) (string, error)
```
Set `isUpdate = false` to create a new user.

### For Update User (Edit Profile)
```go
GenerateImage(isUpdate bool, cwToken string, valueofURL string) (string, error)
```
Set `isUpdate = true` to update a user.

### Example Usage
```go
package main

import (
	"fmt"
	"image-creator/codeToImage/codeToImage"
)

func main() {
	isUpdate := false // Set to true for updating a user
	cwToken := "1234"
	valueofURL := "http://login.cloudtvos.tv/"

	imagePath, err := codeToImage.GenerateImage(isUpdate, cwToken, valueofURL)
	if err != nil {
		fmt.Printf("Error generating image: %v\n", err)
		return
	}

	fmt.Printf("Generated image path: %s\n", imagePath)
}
```

