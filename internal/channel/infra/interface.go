package infra

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DB interface {
	Ping(context.Context) error
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	SendBatch(context.Context, *pgx.Batch) pgx.BatchResults
	BeginTx(context.Context) (pgx.Tx, error)

	//CopyFrom(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) (int64, error)
	//BeginTxOpt(context.Context, pgx.TxOptions) (pgx.Tx, error)
}
