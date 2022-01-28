package endpoints

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"log"
	"math/rand"
	"github.com/streadway/amqp"
	"github.com/gorilla/mux"
)

func failOnError(err error, msg string) {
        if err != nil {
                log.Fatalf("%s: %s", msg, err)
        }
}

func LoadQueryRPC (w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	body := params["num"]
	fmt.Printf("Received a load heavy query, num=%s\n", body)

	conn, err := amqp.Dial("amqp://user:pass@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	fmt.Println("Connected to rabbitmq")

	correlation_id := keyGen()

	routing_key := "load_query_rpc"

	err = ch.Publish(
		"",          // exchange
		routing_key, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: correlation_id,
				ReplyTo:       q.Name,
				Body:          []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	fmt.Printf("Added task to rabbitmq. Key: %s, body: %s\n", routing_key, body)

	for d := range msgs {
			if correlation_id == d.CorrelationId {
					res, err := strconv.Atoi(string(d.Body))
					failOnError(err, "Failed to convert body to integer")

					io.WriteString(w, strconv.Itoa(res) + "\n")
					fmt.Println("Result of arduous calculations", res)

					break
			}
	}
}

func keyGen() string { // generates len 30 key 
	bytes := make([]byte, 30)
	for i := 0; i < 30; i++ {
			bytes[i] = byte(65 + rand.Int31n(26))
	}
	return string(bytes)
}