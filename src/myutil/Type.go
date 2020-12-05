package myutil

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type SublessonInfos struct {
	SublessonInfos []*Sublessoninfo
}

type Sublessoninfo struct {
	SublessonId        int    `json:"sublessonid"`
	SubpackageVersion  string `json:"subpackageversion"`
	BasepackageVersion string `json:"basepackageversion"`
}

type Result struct {
	Sublessoninfo
}

func (r Result) Success() {
	fmt.Println(r.SublessonId)
	ExecuteBuildSublesson(r.Sublessoninfo)
	if r.SublessonId != 0 {
		//构建成功，上传文件
		url := uploadUrl
		fileName := strconv.Itoa(r.SublessonId) + ".zip"
		filePath := buildSublessonFile + fileName
		log.Printf("end execute command,fileName:%s,filePath:%s", fileName, filePath)
		_, err := os.Stat(filePath)
		if err != nil {
			fmt.Println("uoload file not exists")
			return
		}
		UploadFile(url, filePath, fileName, "1")
	}
}


func (r Result) Loading() {

}

func (r Result) Fail() {

}