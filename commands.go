package pgxwrapper

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// Get get SQL command
func (d *DB) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	err := d.db.GetContext(ctx, dest, query, args...)

	return err
}

// GetSq get SQL command with Sqlizer
func (d *DB) GetSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	err = d.db.GetContext(ctx, dest, query, args...)

	return err
}

// Select select SQL command
func (d *DB) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	err := d.db.SelectContext(ctx, dest, query, args...)

	return err
}

// SelectSq select SQL command with Sqlizer
func (d *DB) SelectSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	err = d.db.SelectContext(ctx, dest, query, args...)

	return err
}

// SelectToMapSq select SQL command with Sqlizer with mapper function
func (d *DB) SelectToMapSq(ctx context.Context, dest map[string]interface{}, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	rows, err := d.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.MapScan(dest)
		if err != nil {
			return err
		}
	}

	return err
}

// Insert insert SQL command
func (d *DB) Insert(ctx context.Context, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	_, err := d.db.Exec(query, args...)

	return err
}

// InsertSq insert SQL command with Sqlizer
func (d *DB) InsertSq(ctx context.Context, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	_, err = d.db.ExecContext(ctx, query, args...)

	return err
}

// Delete delete SQL command
func (d *DB) Delete(ctx context.Context, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	err := d.Exec(ctx, query, args...)

	return err
}

// DeleteSq delete SQL command with Sqlizer
func (d *DB) DeleteSq(ctx context.Context, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	err = d.Delete(ctx, query, args...)

	return err
}

// Update update SQL command
func (d *DB) Update(ctx context.Context, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	err := d.Exec(ctx, query, args...)

	return err
}

// UpdateSq update SQL command with Sqlizer
func (d *DB) UpdateSq(ctx context.Context, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	err = d.Update(ctx, query, args...)

	return err
}

// Exec execute SQL command
func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	_, err := d.db.ExecContext(ctx, query, args...)

	return err
}

// ExecSq execute SQL command with Sqlizer
func (d *DB) ExecSq(ctx context.Context, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	err = d.Exec(ctx, query, args...)

	return err
}

// Tx start transaction
func (d *DB) Tx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	if !d.isActive {
		return nil, ErrDBIsNotActive
	}

	return d.db.BeginTxx(ctx, opts)
}

// GetTx get SQL command
func (d *DB) GetTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	err := tx.GetContext(ctx, dest, query, args...)

	return err
}

// GetSqTx get SQL command with Sqlizer
func (d *DB) GetSqTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	err = tx.GetContext(ctx, dest, query, args...)

	return err
}

// SelectTx select SQL command
func (d *DB) SelectTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	err := tx.SelectContext(ctx, dest, query, args...)

	return err
}

// SelectSqTx select SQL command with Sqlizer
func (d *DB) SelectSqTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	err = tx.SelectContext(ctx, dest, query, args...)

	return err
}

// SelectToMapSqTx select SQL command with Sqlizer with mapper function
func (d *DB) SelectToMapSqTx(tx *sqlx.Tx, ctx context.Context, dest map[string]interface{}, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	rows, err := tx.QueryxContext(ctx, query, args...)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.MapScan(dest)
		if err != nil {
			return err
		}
	}

	return err
}

// InsertTx insert SQL command
func (d *DB) InsertTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	_, err := tx.Exec(query, args...)

	return err
}

// InsertSqTx insert SQL command with Sqlizer
func (d *DB) InsertSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)

	return err
}

// DeleteTx delete SQL command
func (d *DB) DeleteTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	_, err := tx.Exec(query, args...)

	return err
}

// DeleteSqTx delete SQL command with Sqlizer
func (d *DB) DeleteSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)

	return err
}

// UpdateTx update SQL command
func (d *DB) UpdateTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	_, err := tx.ExecContext(ctx, query, args...)

	return err
}

// UpdateSqTx update SQL command with Sqlizer
func (d *DB) UpdateSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)

	return err
}

// ExecTx execute SQL command
func (d *DB) ExecTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	_, err := tx.ExecContext(ctx, query, args...)

	return err
}

// ExecSqTx execute SQL command with Sqlizer
func (d *DB) ExecSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)

	return err
}
