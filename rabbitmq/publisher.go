package rabbitmq

import (
	"fmt"
	"log"

	"github.com/micro-search/go-search/domain/entities"
	"github.com/streadway/amqp"
)

type Publisher struct {
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (p *Publisher) Publish(data *entities.Product) {
	conn, err := amqp.Dial("amqp://endrit:endrit@localhost:5672/blla")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"ick.endritqueue", // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)

	failOnError(err, "Failed to declare a queue")

	body := data
	err = ch.Publish(
		"ick.endritexchange", // exchange
		q.Name,               // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%v", body)),
		})
	failOnError(err, "Failed to publish a message")

}
