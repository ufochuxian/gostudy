package myutil

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"tgmnet"
)

const QueueName = "sublessons"

func Send(sublids []*Sublessoninfo, callback tgmnet.Callback) {
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
	failOnError(err, "创建队列失败")
	for i := 0; i < len(sublids); i++ {
		bytes, err := json.Marshal(sublids[i])
		publish(err, ch, q, bytes, callback)
	}
	log.Println("发送消息成功")
}

func publish(err error, ch *amqp.Channel, q amqp.Queue, body []byte, callback tgmnet.Callback) {
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
	}
}
