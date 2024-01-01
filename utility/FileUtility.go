package utility

import (
	"io"
	"mime/multipart"
	"os"
)

type NewFileUtility struct {
}

func (nfu NewFileUtility) Upload(file *multipart.FileHeader, path string) (string, error) {
	src, err := file.Open()

	if err != nil {
		return "", err
	}

	defer src.Close()

	newName := file.Filename

	dst, err := os.Create(path + file.Filename)

	if err != nil {
		return "", err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		return "", err
	}

	return newName, nil
}
