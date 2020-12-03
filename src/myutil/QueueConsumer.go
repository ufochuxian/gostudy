package myutil

import (
	"fmt"
	"github.com/gogf/gf/container/gqueue"
	_ "github.com/gogf/gf/container/gqueue"
	_ "github.com/gogf/gf/os/gtimer"
	"log"
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

	queue := gqueue.New()

	go func() {
		for d := range msgs {
			log.Printf("Received a message : %s", d.Body)
			queue.Push(d.Body)
		}
		fmt.Println("gqueue:")
	}()

	fmt.Println("创建consumer成功")
	//log.Printf(" [*] Waiting for messages, To exit press CTRL+C")
	//<-forever
}
