# pgxwrapper
Simple wrapper for PostgreSQL using pgx, sqlx and squirrel

## Features:
- Compatibility with squirrel
- Transactions methods with squirrel
- Background pinging DB connection with a ticker
- Custom attempts to connect. If attempts is reach their maximum, db connection will close
- Optional receive errors through error channel
- Optional log errors 

## Dependencies:
- github.com/jackc/pgx
- github.com/jmoiron/sqlx
- github.com/Masterminds/squirrel

## Commands (PgDatabase interface):
### Database
1. ```GetDB() *sql.DB```
1. ```IsActive() bool```
1. ```Close() error```
### Without trasactions
1. ```Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error```
1. ```GetSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error```
1. ```Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error```
1. ```SelectSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error```
1. ```SelectToMapSq(ctx context.Context, dest map[string]interface{}, sqlizer sq.Sqlizer) error```
1. ```Insert(ctx context.Context, query string, args ...interface{}) error```
1. ```InsertSq(ctx context.Context, sqlizer sq.Sqlizer) error```
1. ```Delete(ctx context.Context, query string, args ...interface{}) error```
1. ```DeleteSq(ctx context.Context, sqlizer sq.Sqlizer) error```
1. ```Update(ctx context.Context, query string, args ...interface{}) error```
1. ```UpdateSq(ctx context.Context, sqlizer sq.Sqlizer) error```
1. ```Exec(ctx context.Context, query string, args ...interface{}) error```
1. ```ExecSq(ctx context.Context, sqlizer sq.Sqlizer) error```
### With trasactions
1. ```Tx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)```
1. ```GetTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, query string, args ...interface{}) error```
1. ```GetSqTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error```
1. ```SelectTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, query string, args ...interface{}) error```
1. ```SelectSqTx(tx *sqlx.Tx, ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error```
1. ```SelectToMapSqTx(tx *sqlx.Tx, ctx context.Context, dest map[string]interface{}, sqlizer sq.Sqlizer) error```
1. ```InsertTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error```
1. ```InsertSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error```
1. ```DeleteTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error```
1. ```DeleteSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error```
1. ```UpdateTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error```
1. ```UpdateSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error```
1. ```ExecTx(tx *sqlx.Tx, ctx context.Context, query string, args ...interface{}) error```
1. ```ExecSqTx(tx *sqlx.Tx, ctx context.Context, sqlizer sq.Sqlizer) error```

## Options
1. OptionDSN - dsn for connection, example: "host=localhost user=user dbname=db password=123 sslmode=disable"
1. OptionTicker - stands for how often wrapper will check is connection active. Default is 5 sec
1. OptionAttempts - attempts to connect to db. Default is two attempts
1. OptionEnableLogs - enable logging errors and successful attempts to connect using default Golang logger
1. OptionErrorChannel - channel that sends sqlx errors on connection attempts or ErrTooMuchAttempts

## Example
```go
package main

import "github.com/Planck1858/pgxwrapper"

// Optional error channel
errCh := make(chan error)

// Open (with options)
dsn := "host=localhost user=user dbname=db password=password sslmode=disable"
db, err := pgxwrapper.Open(
    pgxwrapper.OptionDSN(dsn),               // standard postgresql DSN
    pgxwrapper.OptionEnableLogs(true))       // use standard Go's log for errors/warnings on connection
    pgxwrapper.OptionTicker(time.Second*10), // how often check db's connection (and reconnect). Default = 5 sec 
    pgxwrapper.OptionAttempts(10),           // attempts to connect to db. Default = 2
    pgxwrapper.OptionErrorChannel(errCh),    // optional error channel that sends errors on connection attempts
)
if err != nil {
	panic(err) // error will cause after first 2 attempts to connect to db or if options are invalid
}

// If you using errCh, then you NEED to receive them
go func() {
	<- errCh
	db.Close()
}()

// Fields should be public!
type user struct {
    Id        string   	`db:"id"`
    Name      string 	`db:"name"`
    CreatedAt time.Time `db:"created_at"`
}

ctx := context.Background()
userId := "eeb6dcb6-0a95-4b30-9b5b-a6e2d09d972b"

// Get
var getId string
err := db.Get(ctx, getId, fmt.Sprintf("SELECT id, name FROM users WHERE id = '%s';", userId))
errCheck(err)

// GetSq
rqGet := sq.Insert("users").
    Columns("id", "name", "created_at").
    Values(UUID.New().String(), "new_user", time.Now().UTC()).
    PlaceholderFormat(sq.Dollar) // For Postgres use only Dollar placeholder

err = db.GetSq(ctx, getId, rqGet)
errCheck(err)

// Select
users := make([]user, 0)
err = db.Select(ctx, &users, "SELECT * FROM users;")
errCheck(err)

// SelectSq
rqSelect := sq.Select("*").From("users")

users := make([]user, 0)
err = db.Select(ctx, &users, rqSelect)
errCheck(err)

// SelectToMapSq
rqSelect := sq.Select("*").From("users").Where(sq.Eq{"id":userId}).PlaceholderFormat(sq.Dollar)

mapUser := make(map[string]interface{})
err = db.SelectToMapSq(ctx, mapUser, rqSelect)
errCheck(err)

userId, ok := mapUser["id"]
if ok {
    userIdStr := userId.(string) 
	...
}

// Insert
err = db.Insert(ctx, "INSERT INTO users (id, name, created_at) VALUES ($1, $2, $3);",
	UUID.New().String(), "new_user1", time.Now().UTC())
errCheck(err)

// InsertSq
rqInsert := sq.Insert("users").
    Columns("id", "name", "created_at").
    Values(UUID.New().String(), "new_user2", time.Now().UTC()).
    PlaceholderFormat(sq.Dollar)

err = db.InserSq(ctx, rqInsert)
errCheck(err)

// Delete
err = db.Delete(ctx, "DELETE FROM users WHERE id = $1;", userId)
errCheck(err)

// DeleteSq
rqDelete := sq.Delete("users").Where(sq.Eq{"id":userId}).PlaceholderFormat(sq.Dollar)

err = db.DeleteSq(ctx, rqDelete)
errCheck(err)

// Update
err = db.Update(ctx, "UPDATE users SET name = $1 WHERE id = $2;", "new_name", userId)
errCheck(err)

// UpdateSq
sqUpdate := sq.Update("users").
    Set("name", "new_name").
    Where(sq.Eq{"id":userId}).
    PlaceholderFormat(sq.Dollar)

err = db.UpdateSq(ctx, sqUpdate)
errCheck(err)

// Exec
err = db.Exec(ctx, "UPDATE users SET name = $1 WHERE id = $2;", "new_name", userId)
errCheck(err)

// ExecSq
err = db.ExecSq(ctx, sqUpdate)
errCheck(err)

// Tx
tx, err := db.Tx(ctx, nil)
errCheck(err)

defer func(){
  if err != nil {
    tx.Rollback()
  } else {
    tx.Commit()
  }
}()

res, err := tx.ExecContext(ctx, "UPDATE users SET name = $1 WHERE id = $2;", "new_name", userId)
errCheck(err)

...
```
