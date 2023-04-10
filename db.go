package pgxwrapper

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type config struct {
	dsn        string
	ticker     time.Duration
	attempts   int
	enableLogs bool
}

type DB struct {
	isActive bool
	db       *sqlx.DB
	ctx      context.Context
	cancel   context.CancelFunc

	config *config
}

func Open(opts ...option) *DB {
	ctx := context.Background()
	c, cf := context.WithCancel(ctx)

	conf := &config{
		dsn:      "",
		ticker:   time.Second,
		attempts: 5,
	}

	for _, o := range opts {
		o(conf)
	}

	db := &DB{
		ctx:    c,
		cancel: cf,
		config: conf,
	}

	db.init()

	return db
}

func (d *DB) GetDB() *sql.DB {
	return d.db.DB
}

func (d *DB) IsActive() bool {
	return d.isActive
}

func (d *DB) init() {
	go func() {
		attempts := 0

		for range time.Tick(d.config.ticker) {
			select {
			case <-d.ctx.Done():
				d.close()
				return

			default:
				switch d.isActive {
				case true:
					if err := d.db.Ping(); err != nil {
						log.Printf("DB failed on ping connection: %s", err)
						d.isActive = false
						break
					}

				default:
					if attempts == d.config.attempts {
						d.close()
						log.Fatalf("DB failed on connect: %s\n", ErrTooMuchAttempts)
						return
					}

					err := d.connect()
					if err != nil {
						log.Printf("DB failed on connect: %s", err)
						attempts++
						break
					}

					d.isActive = true
				}
			}
		}
	}()
}

func (d *DB) connect() error {
	if d.isActive {
		return nil
	}

	db, err := sqlx.Connect("pgx", d.config.dsn)
	if err != nil {
		return err
	}

	d.db = db
	d.isActive = true

	log.Print("success connection to DB")

	return nil
}

func (d *DB) Close() {
	d.close()
	d.cancel()
	log.Print("connection to DB closed")
}

func (d *DB) close() {
	if d.isActive {
		d.db.Close()
		d.isActive = false
	}
}
