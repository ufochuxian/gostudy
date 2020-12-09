package myutil

import (
	"encoding/json"
	"github.com/asaskevich/EventBus"
	_ "github.com/gogf/gf/container/gqueue"
	_ "github.com/gogf/gf/os/gtimer"
	"log"
	"strconv"
)

var globalBus = EventBus.New()

func Receive() {
	log.Println("begin receive")
	globalBus.Subscribe(Pack_Sub_Lesson, func(fileName string) {
		log.Printf("%s，子课程课程构建成功", fileName)
	})
	ch := getChannel()
	q, err := ch.QueueDeclare(
		"sublessons", //name
		false,        //durable
		true,         //delete when unused
		false,        //exclusive
		false,        //no wait
		nil,          //arguments
	)
	failOnError(err, "Failed to declare a que")

	log.Printf("收到消息:%s\n", q.Messages)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	log.Println(msgs)

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

			url := "http://cmscdn.tgmgrp.com/pre_cocos_zip/" + strconv.Itoa(subLid) + ".zip"
			fileName := sublessonSaveFilepath + strconv.Itoa(subLid) + ".zip"

			rs := Result{
				subLessonInfo,
			}
			DownloadFile(url, fileName, rs)
		}
	}()

	log.Println("创建consumer成功")
}
