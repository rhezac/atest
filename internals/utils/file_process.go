package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadDoc(file *multipart.FileHeader, path string) error {
	// Open the uploaded file.
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Create a destination file for the uploaded content.
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the uploaded content to the destination file.
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
