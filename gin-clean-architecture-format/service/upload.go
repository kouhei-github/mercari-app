package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type ImageUrl struct {
	M        string `json:"m"`
	Original string `json:"original"`
}

type BrowserCacheStruct struct {
	Result   bool     `json:"result"`
	ImageId  int      `json:"tmp_img_id"`
	ImageUrl ImageUrl `json:"tmp_img_url"`
}

/** ImageToBrowserCache サーバーに画像を保存する*/
func ImageToBrowserCache(imagePath string, token string, cookie string) (int, error) {
	url := "https://fril.jp/tmp_imgs"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(imagePath)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("image", filepath.Base(imagePath))
	_, errFile1 = io.Copy(part1, file)

	if errFile1 != nil {
		fmt.Println(errFile1)
		return 0, err
	}
	_ = writer.WriteField("authenticity_token", token)
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	req.Header.Add("x-csrf-token", token)
	req.Header.Add("cookie", cookie)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var imageCache BrowserCacheStruct
	err = json.Unmarshal(body, &imageCache)
	if err != nil {
		return 0, err
	}
	if !imageCache.Result {
		err := MyError{Message: "送信はできましたが、画像がラクマにうまく反映されませんでした"}
		return 0, err
	}
	return imageCache.ImageId, nil
}
