package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"time"
)

func StoreImage(imageBlob string) (string, error) {
	// Decode the base64 string
	imageData, err := base64.StdEncoding.DecodeString(imageBlob)
	if err != nil {
		return "", fmt.Errorf("failed to decode image blob: %v", err)
	}

	// Generate a unique file name using the current timestamp
	fileName := fmt.Sprintf("image_%d.jpg", time.Now().UnixNano())
	filePath := filepath.Join("storage/images/fundus", fileName)

	// Ensure the directory exists
	if err := os.MkdirAll("storage/images/fundus", os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Write the image data to the file
	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		return "", fmt.Errorf("failed to write image file: %v", err)
	}

	// Return the file path
	return filePath, nil
}

func ConvertImageToBase64(imagePath string) (string, error) {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %v", err)
	}

	// Create a buffer to hold the image data
	var buf bytes.Buffer

	// Encode the image back to the buffer based on format
	switch format {
	case "jpeg":
		err = jpeg.Encode(&buf, img, nil)
	case "png":
		err = png.Encode(&buf, img)
	default:
		return "", fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return "", fmt.Errorf("failed to encode image: %v", err)
	}

	// Convert the buffer to a base64 string
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64Str, nil
}
