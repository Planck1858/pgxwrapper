package pgxwrapper

import "time"

type option func(c *config)

// OptionDSN dsn for connection, example: "host=localhost:5432 user=user dbname=db password=123 sslmode=disable"
func OptionDSN(dsn string) option {
	return func(c *config) {
		c.dsn = dsn
	}
}

// OptionTicker stands for how often wrapper will check active connection
func OptionTicker(d time.Duration) option {
	return func(c *config) {
		c.ticker = d
	}
}

// OptionAttempts attempts to connect to db
func OptionAttempts(a int) option {
	return func(c *config) {
		c.attempts = a
	}
}

// OptionEnableLogs enable errors and warning log
func OptionEnableLogs(el bool) option {
	return func(c *config) {
		c.enableLogs = el
	}
}

// OptionStartedChannel channel that will return true when connection is active
func OptionStartedChannel(ch chan bool) option {
	return func(c *config) {
		c.startedChannel = ch
	}
}
