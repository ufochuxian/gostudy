package myutil

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"tgmnet"
)

const QueueName = "sublessons"

func Send(sublids []*SubLessonInfo, callback tgmnet.Callback) error {
	log.Println("begin send")
	ch := getChannel()
	q, err := ch.QueueDeclare(
		QueueName, //name
		false,        //durable
		true,         //delete when unused
		false,        //exclusive
		false,        //no wait
		nil,          //arguments
	)
	err = failOnError(err, "创建队列失败")
	if err != nil {
		log.Printf("创建消息队列失败")
		return err
	}
	var publishErr error
	for i := 0; i < len(sublids); i++ {
		bytes, err := json.Marshal(sublids[i])
		publishErr = publish(err, ch, q, bytes, callback)
	}
	if publishErr != nil {
		log.Println("发送打包消息到消息队列失败")
		ReportToFeishu("发送打包消息到消息队列失败", -1)
		return publishErr
	}
	return nil
}

func publish(err error, ch *amqp.Channel, q amqp.Queue, body []byte, callback tgmnet.Callback) error{
	fmt.Printf("publish:%s\n", body)
	err = ch.Publish(
		"",     //exchange
		q.Name, // routing key
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("打包%s失败", body)
		callback(tgmnet.ResponseOld{
			Errno:  100,
			ErrMsg: "打包失败，（发送消息队列失败）",
		})
		return err
	}
	return nil
}
