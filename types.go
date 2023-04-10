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
)

type PgDatabase interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error
	SelectToMapSq(ctx context.Context, cb func(value map[string]interface{}), sqlizer sq.Sqlizer) error
	Insert(ctx context.Context, query string, args ...interface{}) error
	InsertSq(ctx context.Context, sqlizer sq.Sqlizer) error
	Delete(ctx context.Context, query string, args ...interface{}) error
	DeleteSq(ctx context.Context, sqlizer sq.Sqlizer) error
	Update(ctx context.Context, query string, args ...interface{}) error
	UpdateSq(ctx context.Context, sqlizer sq.Sqlizer) error
	Exec(ctx context.Context, query string, args ...interface{}) error
	ExecSq(ctx context.Context, sqlizer sq.Sqlizer) error
	Tx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}
