package db

import (
	"context"
	"fmt"
)

type DB interface {
	Close()
	Reconnect()

	Drop(ctx context.Context)

	InsertRows(ctx context.Context, rows ListRow)

	Get(ctx context.Context, startBlock uint64) ListRow
}

func Init(typeDB TypeDB, nameTable, addr, database, username, password string) DB {
	var db DB

	switch typeDB {
	case CLICKHOUSE:
		db = InitClickHouse(nameTable, addr, database, username, password)
	default:
		panic(fmt.Sprintf("Not correct typeBD = %d", typeDB))
	}

	return db
}
