package myutil

import (
	"bytes"
	_ "github.com/asaskevich/EventBus"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Creates a new file upload http request with optional extra params
func newFileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("X-Token", "sasdsa")
	return request, err
}

func UploadFile(url string, filePath string, fileName string, uoloadFileType string,uploadSubSuccess func()) {
	log.Printf("begin upload,url:%s,filePath:%s,fileName:%s,uoloadFileType:%s", url, filePath, fileName, uoloadFileType)
	extraParams := map[string]string{
		"name": fileName,
		"type": uoloadFileType,
	}
	request, err := newFileUploadRequest(url, extraParams, "body", filePath)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		log.Println(resp.StatusCode)
		log.Println(resp.Header)

		log.Println(body)
	}
	uploadSubSuccess()
	globalBus.Publish(Pack_Sub_Lesson,fileName)
	log.Println("finished upload")
}
