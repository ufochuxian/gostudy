package myutil

import (
	"encoding/json"
	"fmt"
	_ "github.com/gogf/gf/container/gqueue"
	_ "github.com/gogf/gf/os/gtimer"
	"log"
	"strconv"
)

func Receive() {
	fmt.Println("begin receive")
	ch := getChannel()
	q, err := ch.QueueDeclare(
		"sublessons", //name
		false,        //durable
		true,         //delete when unused
		false,        //exclusive
		false,        //no wait
		nil,          //arguments
	)
	failOnError(err, "Failed to declare a queue")

	fmt.Printf("收到消息:%s\n", q.Messages)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	fmt.Println(msgs)

	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message : %s\r", d.Body)
			b := []byte(d.Body)
			subLessonInfo := Sublessoninfo{}
			err := json.Unmarshal(b, &subLessonInfo)
			if err != nil {
				log.Println(err)
			}
			//queue.Push(d.Body)
			subLid := subLessonInfo.SublessonId

			if subLid == 76 {
				url := "http://cmscdn.tgmgrp.com/cocos/zip/" + strconv.Itoa(subLid) + ".zip"
				fileName := sublessonSaveFilepath + strconv.Itoa(subLid) + ".zip"

				rs := Result{
					subLessonInfo,
				}
				DownloadFile(url, fileName, rs)
			}
		}
	}()

	fmt.Println("创建consumer成功")
}
