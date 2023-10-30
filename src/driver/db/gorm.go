package db

import (
	"context"
	"fmt"
	"temp/cmd/config"
	"temp/pkg/logger"
	"temp/pkg/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection interface {
	Connection() (*gorm.DB, error)
	Close(ctx context.Context) error
}

type gormDatabaseConnection struct {
	dsn string

	connectAttempts        int
	connectWaitTimeSeconds int
	connectBlocks          bool
	connecting             bool
	connectError           error

	status          Status
	conn            *gorm.DB
	connectFinished chan bool

	logger logger.Logger
}

func NewConnection(l logger.Logger) Connection {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", config.Config.Cockroach.Host, config.Config.Cockroach.User, config.Config.Cockroach.Password, config.Config.Cockroach.DBName, config.Config.Cockroach.Port, config.Config.Cockroach.SSLMode)

	connectAttempts := config.Config.Cockroach.ConnectAttempts
	if connectAttempts == 0 {
		connectAttempts = 1
	}

	connectWaitTimeSeconds := config.Config.Cockroach.ConnectWaitTime
	if connectWaitTimeSeconds == 0 {
		connectWaitTimeSeconds = 3
	}

	conn := &gormDatabaseConnection{
		dsn:                    dsn,
		connectAttempts:        connectAttempts,
		connectBlocks:          config.Config.Cockroach.ConnectBlocks,
		connectWaitTimeSeconds: connectWaitTimeSeconds,
		status:                 UNKNOWN,
		logger:                 l,
	}

	conn.updateDBStatus()

	return conn
}

func (c *gormDatabaseConnection) Connection() (*gorm.DB, error) {
	c.updateDBStatus()

	if c.status != READY {
		c.tryConnect()
	}

	if c.status == READY {
		return c.conn, nil
	}

	return nil, fmt.Errorf("database connection is not ready: %v, %v", c.status, c.connectError)
}

// tryConnect attempts opening single connection to the database.
func (c *gormDatabaseConnection) tryConnect() {
	if !c.connecting {
		c.connect()
	} else if c.connectBlocks {
		// different goroutine is connecting, wait until finished
		<-c.connectFinished
	}
}

func (c *gormDatabaseConnection) connect() {
	c.connecting = true

	gormConnect := func() (*gorm.DB, error) {
		return gorm.Open(postgres.Open(c.dsn), &gorm.Config{})
	}
	sleep := func(seconds int) {
		time.Sleep(time.Duration(seconds) * time.Second)
	}

	if c.connectAttempts < 0 {
		for c.status != READY {
			c.conn, c.connectError = gormConnect()
			c.updateDBStatus()

			if c.status != READY {
				c.logger.Warnf("unable to connect to database: %v. retrying after %d seconds", c.connectError, c.connectWaitTimeSeconds)
				sleep(c.connectWaitTimeSeconds)
			}
		}

		c.logger.Info("connected with gorm to postgres")
	} else {
		var err error
		for i := 0; i < c.connectAttempts; i++ {
			c.conn, err = gormConnect()
			c.updateDBStatus()

			if c.status != READY {
				c.logger.Warnf("unable to connect to database: %v", err)

				if i < c.connectAttempts-1 {
					sleep(c.connectWaitTimeSeconds)
				}
			} else {
				c.logger.Info("connected with gorm to postgres")
				break
			}
		}

		if c.isConnNil() {
			c.logger.Errorf("failed to connect to database in %d tries: %v", c.connectAttempts, err)
			c.connectError = err
		}
	}

	c.connecting = false
	go func() {
		c.connectFinished <- true
	}()
}

func (c *gormDatabaseConnection) isConnNil() bool {
	return utils.IsInterfaceNil(c.conn)
}

func (c *gormDatabaseConnection) updateDBStatus() {
	if c.isConnNil() {
		c.status = NOT_READY
		return
	}

	sqlDB, err := c.conn.DB()
	if err != nil {
		c.logger.Errorf("failed to get generic SQL database: %v", err)
		c.status = ERROR

		if err = c.Close(context.Background()); err == nil {
			c.status = NOT_READY
		}
		return
	}

	if err := sqlDB.PingContext(context.Background()); err != nil {
		c.logger.Errorf("failed to ping database: %v", err)
		c.status = ERROR

		if err = c.Close(context.Background()); err == nil {
			c.status = NOT_READY
		}
	} else {
		c.status = READY
	}
}

// Close closes the connection to the database.
func (c *gormDatabaseConnection) Close(ctx context.Context) error {
	c.logger.Info("closing gorm postgres db connection")

	if c.isConnNil() {
		c.logger.Debug("no connection to close")
		return nil
	}

	sqlDB, err := c.conn.DB()
	if err != nil {
		c.logger.Errorf("failed to get generic SQL database: %v", err)
		c.status = ERROR
	}

	err = sqlDB.Close()
	if err != nil {
		c.logger.Errorf("failed to close connection: %v", err)
	}

	c.status = DISCONNECTED
	c.conn = nil
	return err
}
