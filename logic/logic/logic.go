package logic

import (
	"logic/async"
	"logic/ram"
)

type Logic struct {
	ram   ram.Ram
	async async.Async
}

func Init(asyncAddr, name, ramAddr string) *Logic {
	ram := ram.Init(ram.REDIS, ramAddr)
	return &Logic{
		ram:   ram,
		async: async.Init(async.RABBITMQ, asyncAddr, name, ram),
	}
}

func (l *Logic) Run() {
	l.async.Receive()
}
