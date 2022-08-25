package service

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadImage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	fileName := strings.ReplaceAll(url, "https://d1gab3f923bew6.cloudfront.net/", "")
	fileName = strings.ReplaceAll(fileName, "JPG", "jpg")
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
	return fileName, nil
}
