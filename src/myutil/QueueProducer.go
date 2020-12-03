package myutil

import (
	"fmt"
	"github.com/streadway/amqp"
)

func Send(sublids []string) {
	fmt.Println("begin send")
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

	for i := 0; i < len(sublids); i++ {
		publish(err, ch, q, sublids[i])
	}
	fmt.Println("发送消息成功")
}

func publish(err error, ch *amqp.Channel, q amqp.Queue, body string) {
	fmt.Printf("publish:%s\n", body)
	err = ch.Publish(
		"",     //exchange
		q.Name, // routing key
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}
