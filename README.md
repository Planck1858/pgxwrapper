# pgxwrapper
Simple wrapper for PostgreSQL using pgx and sqlx + squirrel

Dependencies:
- github.com/jackc/pgx
- github.com/jmoiron/sqlx
- github.com/Masterminds/squirrel

Commands (PgDatabase interface):
- Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
- GetSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error
- Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
- SelectSq(ctx context.Context, dest interface{}, sqlizer sq.Sqlizer) error
- SelectToMapSq(ctx context.Context, dest map[string]interface{}, sqlizer sq.Sqlizer) error
- Insert(ctx context.Context, query string, args ...interface{}) error
- InsertSq(ctx context.Context, sqlizer sq.Sqlizer) error
- Delete(ctx context.Context, query string, args ...interface{}) error
- DeleteSq(ctx context.Context, sqlizer sq.Sqlizer) error
- Update(ctx context.Context, query string, args ...interface{}) error
- UpdateSq(ctx context.Context, sqlizer sq.Sqlizer) error
- Exec(ctx context.Context, query string, args ...interface{}) error
- ExecSq(ctx context.Context, sqlizer sq.Sqlizer) error
- Tx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)

Example:
```go
package main

import "github.com/Planck1858/pgxwrapper"

// Open (with options)
dsn := "host=localhost:5432 user=user dbname=db password=password sslmode=disable"
db := pgxwrapper.Open(
    pgxwrapper.OptionDSN(dsn),              // standard postgresql DSN
    pgxwrapper.OptionTicker(time.Second*5), // how often check db's connection (and reconnect)
    pgxwrapper.OptionAttempts(10),          // attempts to connect to db
    pgxwrapper.OptionEnableLogs(true))      // use standard Go's log for errors/warnings on connection

ctx := context.Background()

// Fields should be public
type user struct {
    Id        string   	`db:"id"`
    Name      string 	`db:"name"`
    CreatedAt time.Time `db:"created_at"`
}

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
