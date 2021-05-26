package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

const QUEUE = "example_task_queue"

func connectRabbitMQ() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/apascualco")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	return conn
}

func openChannel(conn *amqp.Connection) *amqp.Channel {
	c, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	return c
}

func declareQueue(channel *amqp.Channel) amqp.Queue {
	q, err := channel.QueueDeclare(QUEUE, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}
	return q
}

func publish(text string, channel *amqp.Channel, queue amqp.Queue) {
	err := channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(text),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
}

func main() {
	conn := connectRabbitMQ()
	defer conn.Close()

	channel := openChannel(conn)
	defer channel.Close()

	queue := declareQueue(channel)

	publish("apascualco: "+time.Now().String(), channel, queue)

	log.Println("Check queue into: http://localhost:15672/ guest:guest")
}
