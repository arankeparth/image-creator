package codeToImage

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"

	"github.com/chai2010/webp"

	"github.com/golang/freetype"
	"github.com/skip2/go-qrcode"
)

const (
	fontPath           = "./codeToImage/codeToImage/fonts/NotoSans-Regular.ttf"
	editProfileImage   = "./codeToImage/codeToImage/EditProfile.webp"
	createProfileImage = "./codeToImage/codeToImage/CreateProfile.webp"
	fontSize           = 30.0
	qrCodeSize         = 300
	qrPosX             = 1320
	qrPosY             = 610
	textPosX           = 352
	textPosY           = 357
)

func GenerateImage(isUpdate bool, cwToken string, valueofURL string) (string, error) {
	// Parse command-line arguments

	if cwToken == "" || valueofURL == "" {
		fmt.Println("Error: cwToken and valueof_url must be provided.")
		return "", fmt.Errorf("Error: cwToken and valueof_url must be provided.")
	}

	qrCodeLink := valueofURL + cwToken

	// Load the font
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		fmt.Printf("Error reading font file: %v\n", err)
		return "", err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Printf("Error parsing font file: %v\n", err)
		return "", err
	}

	// Load the image (EditProfile.webp or CreateProfile.webp based on isUpdate flag)
	var inputImagePath string
	if isUpdate {
		inputImagePath = editProfileImage
	} else {
		inputImagePath = createProfileImage
	}

	inputFile, err := os.Open(inputImagePath)
	if err != nil {
		fmt.Printf("Error opening image file: %v\n", err)
		return "", err
	}
	defer inputFile.Close()

	img, err := webp.Decode(inputFile)
	if err != nil {
		fmt.Printf("Error decoding webp file: %v\n", err)
		return "", err
	}

	fg := image.NewUniform(color.RGBA{R: 255, G: 255, B: 255, A: 255}) // Text color
	drawer := freetype.NewContext()
	drawer.SetFontSize(fontSize)
	drawer.SetClip(img.Bounds()) // Ensure drawing is within bounds
	drawer.SetDst(img.(draw.Image))
	drawer.SetSrc(fg)
	drawer.SetFont(font)
	// Draw the text at a specific position
	pt := freetype.Pt(textPosX, textPosY+int(drawer.PointToFixed(fontSize)>>6)) // Y position adjusted for font size
	_, err = drawer.DrawString(qrCodeLink, pt)
	if err != nil {
		fmt.Printf("Error drawing text on image: %v\n", err)
		return "", err
	}

	// Generate the QR code
	qr, err := qrcode.New(qrCodeLink, qrcode.Medium)
	if err != nil {
		fmt.Printf("Error generating QR code: %v\n", err)
		return "", err
	}
	qrImage := qr.Image(qrCodeSize) // 300x300 pixels

	// Place the QR code on the image at a specific position
	qrPos := image.Point{X: qrPosX, Y: qrPosY}
	b := qrImage.Bounds()
	rect := image.Rectangle{Min: qrPos, Max: qrPos.Add(b.Size())}
	draw.Draw(img.(draw.Image), rect, qrImage, b.Min, draw.Over)

	// Save the final image
	outputFilePath := fmt.Sprintf("%s.webp", cwToken)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return "", err
	}
	defer outputFile.Close()

	err = webp.Encode(outputFile, img, &webp.Options{Lossless: true})
	if err != nil {
		fmt.Printf("Error saving webp file: %v\n", err)
		return "", err
	}

	return outputFilePath, nil
}
