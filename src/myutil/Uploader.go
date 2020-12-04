package myutil

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func Main() {
	url := "http://dev.cms.tgmgrp.com/v1/resource/cocos/upload"
	FileUpload(url, "./sublesson01.zip")
}

func FileUpload(url string, destFile string) (string, error) {
	fmt.Println("begin FileUpload")
	file, err := os.Open(destFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("X-Token", "d8cdcf8427e")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("finished FileUpload")
	defer response.Body.Close()
	return destFile, nil
}
