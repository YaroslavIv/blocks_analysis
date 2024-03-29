package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type ClickHouse struct {
	conn      driver.Conn
	nameTable string
}

func InitClickHouse(nameTable, addr, database, username, password string) *ClickHouse {
	var (
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
		conn, err   = clickhouse.Open(&clickhouse.Options{
			Addr: []string{addr},
			Auth: clickhouse.Auth{
				Database: database,
				Username: username,
				Password: password,
			},
		})
	)

	defer cancel()

	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		fmt.Printf("Exception %s\n", err)
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		panic(err)
	}

	cl := &ClickHouse{conn: conn, nameTable: nameTable}
	cl.createTable(ctx)
	return cl
}

func (db *ClickHouse) Close() {
	if err := db.conn.Close(); err != nil {
		panic(err)
	}
}

func (db *ClickHouse) Reconnect() {}

func (db *ClickHouse) createTable(ctx context.Context) {
	out, err := db.conn.Query(ctx, fmt.Sprintf("SHOW TABLES FROM default LIKE '%s'", db.nameTable))
	if err != nil {
		panic(err)
	}
	if out.Next() {
		return
	}

	command := fmt.Sprintf(`
	CREATE TABLE %s(
		Hash String NOT NULL,
		Block UInt64 NOT NULL,
		AddrFrom String NOT NULL,
		AddrTo String NOT NULL
	) engine = MergeTree() ORDER BY tuple();`, db.nameTable)

	if err := db.conn.Exec(ctx, command); err != nil {
		panic(err)
	}
}

func (db *ClickHouse) Drop(ctx context.Context) {
	if err := db.conn.Exec(ctx, fmt.Sprintf("DROP TABLE %s", db.nameTable)); err != nil {
		panic(err)
	}
}

func (db *ClickHouse) InsertRows(ctx context.Context, rows ListRow) {
	batch, err := db.conn.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO %s", db.nameTable))
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		if err := batch.Append(row.Data()...); err != nil {
			panic(err)
		}
	}

	if err := batch.Send(); err != nil {
		panic(err)
	}
}

func (db *ClickHouse) Get(ctx context.Context, startBlock uint64) ListRow {
	var rows ListRow

	out, err := db.conn.Query(ctx, fmt.Sprintf("SELECT * from %s WHERE Block >= %d;", db.nameTable, startBlock))
	if err != nil {
		panic(err)
	}

	for out.Next() {
		var (
			hash             string
			block            uint64
			addrFrom, addrTo string
		)

		if err := out.Scan(&hash, &block, &addrFrom, &addrTo); err != nil {
			panic(err)
		}
		rows = append(rows, NewRow(hash, block, addrFrom, addrTo))
	}

	out.Close()
	if err := out.Err(); err != nil {
		panic(err)
	}

	return rows
}
