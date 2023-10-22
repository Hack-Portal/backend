package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/pkg/utils"
	"github.com/jackc/pgx/v5"
)

type Connection interface {
	Connection() (*pgx.Conn, error)
	Close(ctx context.Context) error
}

type connection struct {
	connectError error
	connecting   bool
	status       Status
	conn         *pgx.Conn
	logger       logger.Logger
}

func NewConnection(l logger.Logger) Connection {
	return &connection{logger: l}
}

func (c *connection) Connection() (*pgx.Conn, error) {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.Postgres.ConnectTimeout)*time.Second)
	defer cancel()
	dbUrl := fmt.Sprintf("postgresql://%s:%d/%s?user=%s&password=%s&sslmode=%s",
		config.Config.Postgres.Host,
		config.Config.Postgres.Port,
		config.Config.Postgres.DBName,
		config.Config.Postgres.User,
		config.Config.Postgres.Password,
		config.Config.Postgres.SSLMode,
	)

	pgxConnect := func() (*pgx.Conn, error) {
		return pgx.Connect(ctx, dbUrl)
	}
	sleep := func() {
		time.Sleep(time.Duration(config.Config.Postgres.ConnectWaitTime) * time.Second)
	}

	if config.Config.Postgres.ConnectAttempts == 0 {
		for c.status != READY {
			c.conn, c.connectError = pgxConnect()
			c.updateDBStatus()

			if c.status != READY {
				c.logger.Warnf("unable to connect to database: %v. retrying after %d seconds", c.connectError, config.Config.Postgres.ConnectWaitTime)
			}
		}
		c.logger.Info("connected to database")
	} else {
		var err error

		for i := 0; i < config.Config.Postgres.ConnectAttempts; i++ {
			c.conn, err = pgxConnect()
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

	if err := c.conn.Ping(context.Background()); err != nil {
		c.logger.Errorf("failed to ping database: %v", err)
		c.status = ERROR

		if err = c.Close(context.Background()); err == nil {
			c.status = NOT_READY
		}
	} else {
		c.status = READY
	}
}

func (c *connection) Close(ctx context.Context) error {
	c.logger.Info("closing mongo db connection")

	if c.isConnNil() {
		c.logger.Debug("no connection to close")
		return nil
	}

	err := c.conn.Close(ctx)
	if err != nil {
		c.logger.Errorf("failed to close connection: %v", err)
	}

	c.status = DISCONNECTED
	c.conn = nil
	return err
}
