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

const PackageSublessonUrl = "https://open.feishu.cn/open-apis/bot/v2/hook/4b9d8ca5-fc53-4295-9dc9-dde17d661ea2"

func reportToFeiShu(text string, subId int) {
	if len(text) == 0 {
		fmt.Printf("上报给飞书的字符串不能为空")
		return
	}
		reqStr := buildReq(text, subId)
	req, err := http.NewRequest("POST", PackageSublessonUrl, bytes.NewBufferString(reqStr))
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
//	ReportToFeiShu("haha", 1)
//}
