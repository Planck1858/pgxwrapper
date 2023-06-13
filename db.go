package pgxwrapper

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const twoAttempts = 1

type config struct {
	dsn        string
	ticker     time.Duration
	attempts   int
	enableLogs bool
	errorCh    chan error
}

// DB database main struct
type DB struct {
	isActive bool
	db       *sqlx.DB
	ctx      context.Context
	cancel   context.CancelFunc

	config *config
}

// Open create db entity and start connection to database.
// If connection is not established after attempts on ticker, it will close db connection.
// error will cause after two attempts to connect.
func Open(opts ...option) (*DB, error) {
	ctx := context.Background()
	cCtx, cancel := context.WithCancel(ctx)

	// default config
	conf := &config{
		dsn:        "",
		ticker:     time.Second * 5,
		attempts:   twoAttempts,
		enableLogs: false,
		errorCh:    nil,
	}

	for _, o := range opts {
		o(conf)
	}

	err := validateConfig(conf)
	if err != nil {
		cancel()
		return nil, err
	}

	db := &DB{
		ctx:    cCtx,
		cancel: cancel,
		config: conf,
	}

	startCh := make(chan struct{})
	startErrCh := make(chan error)
	go db.init(startCh, startErrCh)

	select {
	case <-startCh:
	case err := <-startErrCh:
		return nil, err
	}

	return db, nil
}

func validateConfig(conf *config) error {
	if conf.dsn == "" {
		return ErrEmptyDSN
	}

	if conf.attempts < 1 {
		return ErrInvalidAttempts
	}

	return nil
}

// GetDB get sqlx database
func (d *DB) GetDB() *sql.DB {
	return d.db.DB
}

// IsActive get connection status
func (d *DB) IsActive() bool {
	return d.isActive
}

func (d *DB) init(startCh chan struct{}, startErrCh chan error) {
	go func() {
		defer d.close()
		attempts := 0

		for range time.Tick(d.config.ticker) {
			select {
			case <-d.ctx.Done():
				return

			default:
				switch d.isActive {
				case true:
					if err := d.db.Ping(); err != nil {
						if d.isLogEnabled() {
							log.Printf("DB failed to ping connection: %s", err)
						}

						d.isActive = false

						break
					}

				default:
					if attempts >= d.config.attempts {
						if d.isLogEnabled() {
							log.Printf("DB failed to connect after %v attempts: %s",
								d.config.attempts, ErrTooMuchAttempts)
						}

						if !d.isErrChIsNil() {
							d.config.errorCh <- ErrTooMuchAttempts
						}

						return
					}

					err := d.connect()
					if err != nil {
						if d.isLogEnabled() {
							log.Printf("DB failed to connect: %s", err)
						}

						if attempts == twoAttempts {
							startErrCh <- err

							return
						}

						if !d.isErrChIsNil() {
							d.config.errorCh <- err
						}

						attempts++

						break
					}

					if attempts < twoAttempts {
						startCh <- struct{}{}
					}
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

	if d.isLogEnabled() {
		log.Print("successfully connect to DB")
	}

	return nil
}

// Close cancel database connection and cancel context
func (d *DB) Close() {
	d.close()
	if d.isLogEnabled() {
		log.Print("connection to DB is closed")
	}

	d.cancel()
}

func (d *DB) close() {
	if d.isActive {
		err := d.db.Close()
		if err != nil {
			if d.isLogEnabled() {
				log.Printf("error on close DB: %s", err)
			}
		}

		d.isActive = false
	}
}

func (d *DB) isLogEnabled() bool {
	return d.config.enableLogs
}

func (d *DB) isErrChIsNil() bool {
	return d.config.errorCh == nil
}
