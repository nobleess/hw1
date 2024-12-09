package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	dsn  string
	pool *pgxpool.Pool
}

func NewDB(dsn string) *DB {
	return &DB{
		dsn:  dsn,
		pool: nil,
	}
}

func (db *DB) Open(ctx context.Context) (err error) {
	db.pool, err = NewPool(ctx, db.dsn)
	return err
}

func NewPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pgxPoolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}
	return pgxpool.NewWithConfig(ctx, pgxPoolConfig)
}

func (db *DB) Close() error {
	if db.pool != nil {
		db.pool.Close()
	}
	return nil
}

func (db *DB) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

func (db *DB) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return db.pool.Query(ctx, sql, args...)
}

func (db *DB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return db.pool.QueryRow(ctx, sql, args...)
}

func (db *DB) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, sql, args...)
}

func (db *DB) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return db.pool.Begin(ctx)
}

func (db *DB) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return db.pool.SendBatch(ctx, b)
}
