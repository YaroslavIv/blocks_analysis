package async

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel

	msgs <-chan amqp.Delivery

	answers map[string](chan []byte)

	name string
}

func InitRabbitMQ(url, name string) *RabbitMQ {
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil || ch == nil {
		panic(err)
	}

	_, err = ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)

	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		"amq.rabbitmq.reply-to", // queue
		"ReplyToConsumer",       // consumer
		true,                    // auto-ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-wait
		nil,                     // args
	)
	if err != nil {
		panic(err)
	}

	r := &RabbitMQ{
		conn:    conn,
		ch:      ch,
		name:    name,
		msgs:    msgs,
		answers: make(map[string](chan []byte)),
	}

	go r.receive()

	return r
}

func (a *RabbitMQ) Send() string {
	CorrelationId := strconv.Itoa(rand.Intn(9999999999))
	fmt.Printf("Send: %s\n", CorrelationId)

	err := a.ch.Publish(
		"",
		a.name,
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: CorrelationId,
			Body:          []byte("TOP5"),
			ReplyTo:       "amq.rabbitmq.reply-to",
		},
	)
	if err != nil {
		panic(err)
	}
	a.answers[CorrelationId] = make(chan []byte)

	return CorrelationId
}

func (a *RabbitMQ) receive() {
	for {
		select {
		case msg := <-a.msgs:
			a.answers[msg.CorrelationId] <- msg.Body
		}
	}
}

func (a *RabbitMQ) Get(query string) Top {
	var out Top

	if err := json.Unmarshal(<-a.answers[query], &out); err != nil {
		panic(err)
	}

	return out
}
