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

func (r *Redis) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *Redis) GetBlockNumber(ctx context.Context) uint64 {
	var out uint64
	r.get(ctx, "BlockNumber", &out)

	return out
}

func (r *Redis) SetBlockNumber(ctx context.Context, value uint64) error {
	return r.set(ctx, "BlockNumber", value)
}

func (r *Redis) GetTop(ctx context.Context, block uint64) Top {
	var out Top
	r.get(ctx, fmt.Sprintf("Top%d", block), &out)

	return out
}

func (r *Redis) SetTop(ctx context.Context, block uint64, value Top) error {
	return r.set(ctx, fmt.Sprintf("Top%d", block), value)
}

func (r *Redis) GetERC20(ctx context.Context) ListERC20 {
	var out ListERC20
	r.get(ctx, "ERC20", &out)

	return out
}

func (r *Redis) SetERC20(ctx context.Context, value ListERC20) error {
	return r.set(ctx, "ERC20", value)
}

func (r *Redis) get(ctx context.Context, key string, value interface{}) {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &value)
	if err != nil {
		panic(err)
	}
}

func (r *Redis) set(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return r.client.Set(ctx, key, data, 0).Err()
}
