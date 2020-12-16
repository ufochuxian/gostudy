package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// get current project's root path
// return path not contain the exec file
func GetProjectRoot() string {
	var (
		path string
		err  error
	)
	defer func() {
		if err != nil {
			panic(fmt.Sprintf("GetProjectRoot error :%+v", err))
		}
	}()
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}

type Content struct {
	Text string `json:"text"`
}

type ReportRequest struct {
	MsgType string  `json:"msg_type"`
	Content Content `json:"content"`
}

const PACKAGE_SUBLESSON_URL = "https://open.feishu.cn/open-apis/bot/v2/hook/4b9d8ca5-fc53-4295-9dc9-dde17d661ea2"

func ReportToFeishu(text string, subId int) {
	reqStr := buildReq(text, subId)
	req, err := http.NewRequest("POST", PACKAGE_SUBLESSON_URL, bytes.NewBufferString(reqStr))
	if err != nil {
		fmt.Printf(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer resp.Body.Close()
}

func buildReq(text string, subId int) string {
	respMsg := "子课程:" + strconv.Itoa(subId) + "进行打包，返回信息:" + text
	req := ReportRequest{
		MsgType: "text",
		Content: Content{
			Text: respMsg,
		},
	}
	marshal, err := json.Marshal(req)
	if err == nil {
		fmt.Printf(string(marshal))
	}
	return string(marshal)
}

//func main() {
//	//log.Println(GetProjectRoot())
//	ReportToFeishu("haha", 1)
//}
