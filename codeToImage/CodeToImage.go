package codeToImage

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"

	"github.com/golang/freetype"
	"github.com/skip2/go-qrcode"
)

const (
	fontPath           = "./codeToImage/fonts/NotoSans-Regular.ttf"
	editProfileImage   = "./codeToImage/EditProfile.png"
	createProfileImage = "./codeToImage/CreateProfile.png"
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

	var wg sync.WaitGroup
	var qrImage image.Image
	var qrErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Generate the QR code
		qr, err := qrcode.New(qrCodeLink, qrcode.Medium)
		if err != nil {
			qrErr = err
			return
		}
		qrImage = qr.Image(qrCodeSize) // 300x300 pixels
	}()

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

	img, err := png.Decode(inputFile)
	if err != nil {
		fmt.Printf("Error decoding png file: %v\n", err)
		return "", err
	}

	fg := image.NewUniform(color.RGBA{R: 4, G: 186, B: 250, A: 255}) // Text color
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

	wg.Wait()
	if qrErr != nil {
		fmt.Printf("Error generating QR code: %v\n", qrErr)
		return "", qrErr
	}

	// Place the QR code on the image at a specific position
	qrPos := image.Point{X: qrPosX, Y: qrPosY}
	b := qrImage.Bounds()
	rect := image.Rectangle{Min: qrPos, Max: qrPos.Add(b.Size())}
	draw.Draw(img.(draw.Image), rect, qrImage, b.Min, draw.Over)

	// Encode the final image to PNG format
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		fmt.Printf("Error encoding png file: %v\n", err)
		return "", err
	}

	// Convert the encoded image to a Base64 string
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64Str, nil
}
