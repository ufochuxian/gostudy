package myutil

import "github.com/streadway/amqp"

// create connect
func openConnect() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()
	return conn
}

func getChannel() *amqp.Channel {
	conn := openConnect()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()
	return ch
}
