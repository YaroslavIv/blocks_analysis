package ram

import "fmt"

type Ram interface {
	Del(key string) error

	GetBlockNumber() uint64
	SetBlockNumber(value uint64) error

	GetTop(block uint64) Top
	SetTop(block uint64, top Top) error

	GetERC20() ListERC20
	SetERC20(value ListERC20) error
}

func Init(typeRam TypeRAM, addr string) Ram {
	var ram Ram

	switch typeRam {
	case REDIS:
		ram = InitRedis(addr)
	default:
		panic(fmt.Sprintf("Not correct typeRam = %d", typeRam))
	}

	return ram
}
