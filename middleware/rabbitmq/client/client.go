package client

import (
	"github.com/streadway/amqp"
	"github.com/jhgv/gocodes/middleware/utils/constants"
	"github.com/jhgv/gocodes/middleware/utils"
	"log"
)

func StartClient() {
	conn, err := amqp.Dial(constants.RabbitMQHost)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	requestQueue, err := ch.QueueDeclare(constants.RequestQueue, false, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare the request queue")

	responseQueue, err := ch.QueueDeclare(constants.ResponseQueue, false, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare the response queue")


	//startTimes := make(chan time.Time)
	//elapsedTimes := make(chan time.Duration)
	responses, err := ch.Consume(
		responseQueue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	//go func() {
	//	for t := range times {
	//		elapsedReq := time.Since(t)
	//		log.Println(elapsedReq.String())
	//	}
	//}()

	count := 0
	for ; count < constants.NumRepetitions; {
		body := utils.GenerateRandomText(20)
		//times <- time.Now()
		ch.Publish(
			"",     // exchange
			requestQueue.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing {
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		count ++
	}

	for r := range responses {
		//elapsedTimes <-
		log.Printf("[x] Message recieved from server: %s", r.Body)
	}

}