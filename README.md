# pgxwrapper
Simple wrapper for PostgreSQL using pgx and sqlx

Example:
```go
package main

import "github.com/Planck1858/pgxwrapper"

db := pgxwrapper.Open(
    pgxwrapper.OptionDSN("host=localhost:5432 user=user dbname=db password=password sslmode=disable"), // standard postgresql DSN
    pgxwrapper.OptionTicker(time.Second*5), // how often check db's connection (and reconnect)
    pgxwrapper.OptionAttempts(10),          // attempts to connect to db
    pgxwrapper.OptionEnableLogs(true))      // use standard Go's log for errors/warnings on connection

ctx := context.Background()

// Fields should be public
type user struct {
    Id string   `db:"id"`
	Name string `db:"name"`
}

// Get
var getId string
err := db.Get(ctx, getId, "SELECT id, name FROM users WHERE ", nil)
errCheck(err)

// Select
users := make([]user, 0)
err = db.Select(ctx, &users, "SELECT * FROM users")
errCheck(err)

// InsertSq
rqi := sq.Insert("users").
    Columns("id", "name", "created_at").
    Values(UUID.New().String(), "new_user", time.Now().UTC()).
    PlaceholderFormat(sq.Dollar) // For Postgres use only Dollar placeholder

err = db.InserSq(ctx, rqi)
errCheck(err)

// TODO: other commands
```

Dependencies: 
- github.com/jackc/pgx
- github.com/jmoiron/sqlx
- github.com/Masterminds/squirrel
