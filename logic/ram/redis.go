package ram

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func InitRedis(ramAddr string) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     ramAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Redis{client: client}
}

func (r *Redis) Del(key string) error {
	return r.client.Del(context.Background(), key).Err()
}

func (r *Redis) GetBlockNumber() uint64 {
	var out uint64
	r.get("BlockNumber", &out)

	return out
}

func (r *Redis) SetBlockNumber(value uint64) error {
	return r.set("BlockNumber", value)
}

func (r *Redis) GetTop(block uint64) Top {
	var out Top
	r.get(fmt.Sprintf("Top%d", block), &out)

	return out
}

func (r *Redis) SetTop(block uint64, value Top) error {
	return r.set(fmt.Sprintf("Top%d", block), value)
}

func (r *Redis) GetERC20() ListERC20 {
	var out ListERC20
	r.get("ERC20", &out)

	return out
}

func (r *Redis) SetERC20(value ListERC20) error {
	return r.set("ERC20", value)
}

func (r *Redis) get(key string, value interface{}) {
	data, err := r.client.Get(context.Background(), key).Bytes()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &value)
	if err != nil {
		panic(err)
	}
}

func (r *Redis) set(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return r.client.Set(context.Background(), key, data, 0).Err()
}
