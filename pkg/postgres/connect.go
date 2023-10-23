package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/pkg/utils"
)

type Connection interface {
	Connection() (*sql.DB, error)
	Close(ctx context.Context)
}

type connection struct {
	connectError error
	connecting   bool
	status       Status
	conn         *sql.DB
	logger       logger.Logger
}

func NewConnection(l logger.Logger) Connection {
	return &connection{logger: l}
}

func (c *connection) Connection() (*sql.DB, error) {
	c.updateDBStatus()

	if c.status != READY {
		c.tryConnect()
	}

	if c.status == READY {
		return c.conn, nil
	}
	return nil, fmt.Errorf("database connection is not ready: %v, %v", c.status, c.connectError)
}

func (c *connection) tryConnect() {
	if !c.connecting {
		c.connect()
	}
}

func (c *connection) connect() {
	c.connecting = true
	dbUrl := fmt.Sprintf("postgresql://%s:%d/%s?user=%s&password=%s&sslmode=%s",
		config.Config.Postgres.Host,
		config.Config.Postgres.Port,
		config.Config.Postgres.DBName,
		config.Config.Postgres.User,
		config.Config.Postgres.Password,
		config.Config.Postgres.SSLMode,
	)

	dbConnect := func() (*sql.DB, error) {
		return sql.Open("postgres", dbUrl)
	}
	sleep := func() {
		time.Sleep(time.Duration(config.Config.Postgres.ConnectWaitTime) * time.Second)
	}

	if config.Config.Postgres.ConnectAttempts == 0 {
		for c.status != READY {
			c.conn, c.connectError = dbConnect()
			c.updateDBStatus()

			if c.status != READY {
				c.logger.Warnf("unable to connect to database: %v. retrying after %d seconds", c.connectError, config.Config.Postgres.ConnectWaitTime)
			}
		}
		c.logger.Info("connected to database")
	} else {
		var err error

		for i := 0; i < config.Config.Postgres.ConnectAttempts; i++ {
			c.conn, err = dbConnect()
			c.updateDBStatus()

			if c.status != READY {
				c.logger.Warnf("unable to connect to database: %v", err)

				if i < config.Config.Postgres.ConnectAttempts-1 {
					sleep()
				}
			} else {
				c.logger.Info("connected to database")
				break
			}
		}

		if c.isConnNil() {
			c.logger.Errorf("failed to connect to database in %d tries: %v", config.Config.Postgres.ConnectAttempts, err)
			c.connectError = err
		}
	}

	c.connecting = false
}

func (c *connection) isConnNil() bool {
	return utils.IsInterfaceNil(c.conn)
}

func (c *connection) updateDBStatus() {
	if c.isConnNil() {
		c.status = NOT_READY
		return
	}

	if err := c.conn.Ping(); err != nil {
		c.logger.Errorf("failed to ping database: %v", err)
		c.status = ERROR

		c.Close(context.Background())
	} else {
		c.status = READY
	}
}

func (c *connection) Close(ctx context.Context) {
	c.logger.Info("closing mongo db connection")

	if c.isConnNil() {
		c.logger.Debug("no connection to close")
		return
	}

	c.conn.Close()

	c.status = DISCONNECTED
	c.conn = nil
}
