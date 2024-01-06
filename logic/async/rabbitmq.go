package async

import (
	"context"
	"encoding/json"
	"logic/ram"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	ram  ram.Ram

	msgs <-chan amqp.Delivery

	name string
}

func InitRabbitMQ(url, name string, ram ram.Ram) *RabbitMQ {
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
		name,              // queue
		"ReplyToConsumer", // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		panic(err)
	}

	r := &RabbitMQ{
		conn: conn,
		ch:   ch,
		name: name,
		msgs: msgs,
		ram:  ram,
	}

	return r
}
func (a *RabbitMQ) Receive(ctx context.Context) error {
	for {
		select {
		case msg := <-a.msgs:

			data, err := json.Marshal(a.ram.GetTop(ctx, a.ram.GetBlockNumber(ctx)))
			if err != nil {
				return err
			}
			err = a.ch.PublishWithContext(
				ctx,
				"",          // exchange
				msg.ReplyTo, // routing key
				false,       //
				false,       // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: msg.CorrelationId,
					Body:          data,
				})
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
