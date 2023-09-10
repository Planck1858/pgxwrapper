package pgxwrapper

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	ErrDBIsNotActive   = errors.New("connect not active")
	ErrTooMuchAttempts = errors.New("database closed: too much attempts")
	ErrEmptyDSN        = errors.New("empty dsn")
	ErrInvalidAttempts = errors.New("invalid attempts number, must be > 0")
)

type PgDatabase interface {
	GetDB() *sql.DB
	IsActive() bool

	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error
	SelectToMapSq(ctx context.Context, dest map[string]interface{}, sqlizer sq.Sqlizer) error
	Insert(ctx context.Context, query string, args ...interface{}) error
	InsertSq(ctx context.Context, sqlizer sq.Sqlizer) error
	Delete(ctx context.Context, query string, args ...interface{}) error
	DeleteSq(ctx context.Context, sqlizer sq.Sqlizer) error
	Update(ctx context.Context, query string, args ...interface{}) error
	UpdateSq(ctx context.Context, sqlizer sq.Sqlizer) error
	Exec(ctx context.Context, query string, args ...interface{}) error
	ExecSq(ctx context.Context, sqlizer sq.Sqlizer) error

	Tx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	GetTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetSqTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error
	SelectTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectSqTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error
	SelectToMapSqTx(tx *sqlx.Tx, ctx context.Context, dest map[string]interface{}, sqlizer sq.Sqlizer) error
	InsertTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error
	InsertSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error
	DeleteTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error
	DeleteSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error
	UpdateTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error
	UpdateSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error
	ExecTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error
	ExecSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error
}
