package async

import "fmt"

type Async interface {
	Send() string
	Get(query string) Top
}

func Init(typeAsync TypeAsync, url, name string) Async {
	var async Async

	switch typeAsync {
	case RABBITMQ:
		async = InitRabbitMQ(url, name)
	default:
		panic(fmt.Sprintf("Not correct typeAsync = %d", typeAsync))
	}

	return async
}
