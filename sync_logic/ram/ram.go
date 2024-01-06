package ram

import (
	"context"
	"fmt"
)

type Ram interface {
	Del(ctx context.Context, key string) error

	GetBlockNumber(ctx context.Context) uint64
	SetBlockNumber(ctx context.Context, value uint64) error

	GetTop(ctx context.Context, block uint64) Top
	SetTop(ctx context.Context, block uint64, top Top) error

	GetERC20(ctx context.Context) ListERC20
	SetERC20(ctx context.Context, value ListERC20) error
}

func Init(typeRam TypeRAM, ramAddr string) Ram {
	var ram Ram

	switch typeRam {
	case REDIS:
		ram = InitRedis(ramAddr)
	default:
		panic(fmt.Sprintf("Not correct typeRam = %d", typeRam))
	}

	return ram
}
