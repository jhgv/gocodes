package server

import (
	"github.com/streadway/amqp"
	"github.com/jhgv/gocodes/middleware/utils/constants"
	"github.com/jhgv/gocodes/middleware/utils"
	"log"
)

func StartServer() {
	conn, err := amqp.Dial(constants.RabbitMQHost)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	requestQueue, err := ch.QueueDeclare(constants.RequestQueue,false,false,false,false,nil)
	utils.FailOnError(err, "Failed to declare a queue")

	responseQueue, err := ch.QueueDeclare(constants.ResponseQueue,false,false,false,false,nil)
	utils.FailOnError(err, "Failed to declare a queue")

	requests, err := ch.Consume(
		requestQueue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	count := 1
	for d := range requests {
		log.Printf("Received a message: %s", d.Body)
		res := utils.ProcessedMessage(string(d.Body))
		err = ch.Publish(
			"",     // exchange
			responseQueue.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing {
				ContentType: "text/plain",
				Body:        []byte(res),
			})
		count++
		if count == constants.NumRepetitions {
			//forever <- false
		}
	}
}