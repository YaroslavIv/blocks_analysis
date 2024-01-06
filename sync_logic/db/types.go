package db

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

type TypeDB int

const (
	CLICKHOUSE TypeDB = iota
)

type Row struct {
	Hash     string
	Block    uint64
	AddrFrom string
	AddrTo   string
}

func NewRow(hash string, block uint64, addrFrom, addrTo string) Row {
	return Row{
		Hash:     hash,
		Block:    block,
		AddrFrom: addrFrom,
		AddrTo:   addrTo,
	}
}

func (r Row) Data() []interface{} {
	return []interface{}{r.Hash, r.Block, r.AddrFrom, r.AddrTo}
}

func (r Row) String() string {
	out := fmt.Sprintf("Hash: %s\n", r.Hash) +
		fmt.Sprintf("Block: %d\n", r.Block) +
		fmt.Sprintf("AddrFrom: %s\n", r.AddrFrom) +
		fmt.Sprintf("AddrTo: %s\n", r.AddrTo)

	return out
}

type ListRow []Row

func NewListRow(block *types.Block) ListRow {
	var rows ListRow
	txs := block.Transactions()
	for _, tx := range txs {
		var to string
		if tx.To() == nil {
			to = ""
		} else {
			to = tx.To().Hex()
		}
		rows = append(rows, NewRow(tx.Hash().Hex(), block.NumberU64(), "", to))
	}

	return rows
}

func (lr ListRow) Data() []interface{} {
	var out []interface{}

	for _, row := range lr {
		out = append(out, row.Data()...)
	}

	return out
}
