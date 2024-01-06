package erc20

import (
	"fmt"
	"reflect"
	"strings"
	"sync_logic/db"

	"github.com/ethereum/go-ethereum/core/types"
)

type ERC20 struct {
}

func Init() *ERC20 {
	return &ERC20{}
}

var SignTransfer = []byte{169, 5, 156, 187}

const TransferLen = 68

func (e *ERC20) Check(block uint64, txs types.Transactions) db.ListRow {
	var rows db.ListRow

	for _, tx := range txs {

		if tx.To() == nil || len(tx.Data()) != TransferLen || !reflect.DeepEqual(tx.Data()[:4], SignTransfer) {
			continue
		}

		addr, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			panic(err)
		}
		addrFrom := strings.ToLower(addr.Hex())
		addrTo := strings.ToLower(fmt.Sprintf("0x%x", tx.Data()[16:36]))

		rows = append(rows, db.NewRow(tx.Hash().Hex(), block, addrFrom, addrTo))
	}

	return rows
}
