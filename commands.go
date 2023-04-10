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

// Getx get SQL command with Sqlizer
func (d *DB) Getx(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error {
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

// Selectx select SQL command with Sqlizer
func (d *DB) Selectx(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error {
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

// SelectToMap select SQL command with Sqlizer with map destination
func (d *DB) SelectToMap(ctx context.Context, cb func(value map[string]interface{}), sqlizer sq.Sqlizer) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	rows, err := d.db.Queryx(query, args...)
	if err != nil {
		return err
	}

	for rows.Next() {
		results := make(map[string]interface{}, 1)
		err = rows.MapScan(results)
		if err != nil {
			return err
		}
		if cb != nil {
			cb(results)
		}
	}

	return err
}

// Insert inser SQL command
func (d *DB) Insert(ctx context.Context, query string, args ...interface{}) error {
	if !d.isActive {
		return ErrDBIsNotActive
	}

	_, err := d.db.Exec(query, args...)

	return err
}

// Insertx insert SQL command with Sqlizer
func (d *DB) Insertx(ctx context.Context, sqlizer sq.Sqlizer) error {
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

// Deletex delete SQL command with Sqlizer
func (d *DB) Deletex(ctx context.Context, sqlizer sq.Sqlizer) error {
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

// Updatex update SQL command with Sqlizer
func (d *DB) Updatex(ctx context.Context, sqlizer sq.Sqlizer) error {
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

// Execx execute SQL command with Sqlizer
func (d *DB) Execx(ctx context.Context, sqlizer sq.Sqlizer) error {
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
