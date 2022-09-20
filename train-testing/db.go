//go:generate mockgen -source=$GOFILE -destination mock/mock_$GOFILE -package=mock
package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB interface {
	Ping(ctx context.Context) error
}

type db struct {
	db *sql.DB
}

var _ DB = (*db)(nil)

func New(d *sql.DB) DB {
	return &db{
		db: d,
	}
}

type Usecase func(db DB, ctx context.Context) error

func NewUsecase() Usecase {
	return func(db DB, ctx context.Context) error {
		return db.Ping(ctx)
	}
}

func (d *db) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}
