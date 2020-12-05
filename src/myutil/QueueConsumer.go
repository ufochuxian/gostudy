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

	//forever := make(chan bool)

	//queue := gqueue.New()

	go func() {
		for d := range msgs {
			log.Printf("Received a message : %s\r", d.Body)
			b := []byte(d.Body)
			sublessoninfo := Sublessoninfo{}
			err := json.Unmarshal(b, &sublessoninfo)
			if err != nil {
				fmt.Println(err)
			}
			//queue.Push(d.Body)
			sublid := sublessoninfo.SublessonId

			if sublid == 1 {
				url := "http://cmscdn.tgmgrp.com/cocos/zip/" + strconv.Itoa(sublid) + ".zip"
				fileName := sublessonSaveFilepath + strconv.Itoa(sublid) + ".zip"

				rs := Result{
					sublessoninfo,
				}
				DownloadFile(url, fileName,rs)
			}
		}
	}()

	fmt.Println("创建consumer成功")
	//log.Printf(" [*] Waiting for messages, To exit press CTRL+C")
	//<-forever
}


