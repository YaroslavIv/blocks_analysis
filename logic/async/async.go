package async

import (
	"context"
	"fmt"
	"logic/ram"
)

type Async interface {
	Receive(ctx context.Context) error
}

func Init(typeAsync TypeAsync, url, name string, ram ram.Ram) Async {
	var async Async

	switch typeAsync {
	case RABBITMQ:
		async = InitRabbitMQ(url, name, ram)
	default:
		panic(fmt.Sprintf("Not correct typeAsync = %d", typeAsync))
	}

	return async
}
