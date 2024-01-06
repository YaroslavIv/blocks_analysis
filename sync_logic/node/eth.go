package node

import (
	"context"
	"fmt"
	"sync_logic/db"
	"sync_logic/erc20"
	"sync_logic/ram"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Eth struct {
	client        *ethclient.Client
	header        chan *types.Header
	sub           ethereum.Subscription
	db            db.DB
	checker       *erc20.ERC20
	ram           ram.Ram
	maxCountBlock uint64
}

func InitEth(rawurl string, maxCountBlock uint64, nameTable, ramAddr, dbAddr, dbDatabase, dbUsername, dbPassword string) *Eth {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		panic(err)
	}

	eth := &Eth{
		client:        client,
		header:        make(chan *types.Header),
		db:            db.Init(db.CLICKHOUSE, nameTable, dbAddr, dbDatabase, dbUsername, dbPassword),
		checker:       erc20.Init(),
		ram:           ram.Init(ram.REDIS, ramAddr),
		maxCountBlock: maxCountBlock,
	}

	eth.subscribe()

	return eth
}

func (eth *Eth) subscribe() {
	sub, err := eth.client.SubscribeNewHead(context.Background(), eth.header)
	if err != nil {
		panic(err)
	}
	eth.sub = sub
}

func (eth *Eth) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {
		case err := <-eth.sub.Err():
			return err
		case header := <-eth.header:
			hash := header.Hash()
			block := eth.BlockByHash(ctx, hash)
			number := block.NumberU64()

			rowsNew := eth.checker.Check(number, block.Transactions())
			rowsOld := eth.db.Get(ctx, number-eth.maxCountBlock+1)
			top := getTopFive(rowsNew, rowsOld)

			eth.ram.SetTop(ctx, number, top)
			eth.ram.SetBlockNumber(ctx, number)
			eth.db.InsertRows(ctx, rowsNew)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (eth *Eth) BlockByHash(ctx context.Context, hash common.Hash) *types.Block {
	for {
		block, err := eth.client.BlockByHash(ctx, hash)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			time.Sleep(time.Millisecond * 100)
		} else {
			return block
		}
	}
}
