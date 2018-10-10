package client

import (
	"fmt"
	"log"
	"time"

	"github.com/jhgv/gocodes/middleware/utils"
	"github.com/jhgv/gocodes/middleware/utils/constants"
	"github.com/streadway/amqp"
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

	// Xlsx file operations
	xlsBuilder := utils.XlsxBuilder{}
	fileName := fmt.Sprintf("rabbitmq-%d.xlsx", constants.NumRepetitions)
	xlsBuilder.SetFileName(fileName)
	xlsBuilder.CreateHeader()
	averageFormula := fmt.Sprintf("AVERAGE(A%d:A%d)", xlsBuilder.GetRowNum()+1, constants.NumRepetitions+2)
	xlsBuilder.SetupAverageFormula(averageFormula)

	startTimes := make(chan time.Time, constants.NumRepetitions)
	elapsedTimes := make(chan time.Duration, constants.NumRepetitions)
	responses, err := ch.Consume(
		responseQueue.Name, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)

	go func() {
		for t := range elapsedTimes {
			xlsBuilder.AddRowData(t.Seconds() * 1000.0)
		}
	}()

	count := 0
	for count < constants.NumRepetitions {
		body := utils.GenerateRandomText(20)

		ch.Publish(
			"",                // exchange
			requestQueue.Name, // routing key
			false,             // mandatory
			false,             // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		startTimes <- time.Now()
		count++
	}

	for r := range responses {
		startTime := <-startTimes
		log.Print(len(startTimes))
		elapsedTimes <- time.Since(startTime)
		log.Printf("Message from server: %s", r.Body)
		xlsBuilder.GenerateFile()
	}
}
