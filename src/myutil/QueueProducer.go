package myutil

import (
	"fmt"
	"github.com/streadway/amqp"
)

func Send() {
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
	body := "hello"

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
	fmt.Println("创建producer成功")
}
