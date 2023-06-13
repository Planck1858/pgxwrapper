package pgxwrapper

import "time"

type option func(c *config)

// OptionDSN dsn for connection, example: "host=localhost user=user dbname=db password=123 sslmode=disable"
func OptionDSN(dsn string) option {
	return func(c *config) {
		c.dsn = dsn
	}
}

// OptionTicker stands for how often wrapper will check is connection active
func OptionTicker(d time.Duration) option {
	return func(c *config) {
		c.ticker = d
	}
}

// OptionAttempts attempts to connect to db. Default is two attempts
func OptionAttempts(a int) option {
	return func(c *config) {
		c.attempts = a
	}
}

// OptionEnableLogs enable logging errors and successful attempts to connect using default Golang logger
func OptionEnableLogs(el bool) option {
	return func(c *config) {
		c.enableLogs = el
	}
}

// OptionErrorChannel channel that sends sqlx errors on connection attempts or ErrTooMuchAttempts
func OptionErrorChannel(ch chan error) option {
	return func(c *config) {
		c.errorCh = ch
	}
}
