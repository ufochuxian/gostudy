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
		UploadFile(url, filePath, fileName, "1", func() {
			//上传成功
			log.Println("上传子课程成功")
			mainFileName := "Main.zip"
			mainFilePath := buildSublessonFile + mainFileName
			log.Printf("开始上传Main.zip,mainFileName:%s,mainFilePath:%s", mainFileName, mainFilePath)
			_, err := os.Stat(mainFilePath)
			if err != nil {
				log.Println("mainFile not exists")
				return
			}
			UploadFile(url, mainFilePath, mainFileName, "1", func() {
				//上传成功
				log.Printf("上传main成功")
			})
		})
	}
}

func (r Result) Loading() {

}

func (r Result) Fail() {

}
