package gorm

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/hackhack-Geek-vol6/backend/src/frameworks/db"
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

	status          db.Status
	conn            *gorm.DB
	connectFinished chan bool

	logger *slog.Logger
}

func NewGormConnection(
	dsn string,
	connectAttempts int,
	connectWaitTimeSeconds int,
	connectBlocks bool,

	logger *slog.Logger,
) Connection {
	if logger == nil {
		logger = slog.Default().With(slog.String("package", "db"), slog.String("framework", "gorm"))
	}

	if connectAttempts == 0 {
		connectAttempts = 1
	}

	if connectWaitTimeSeconds == 0 {
		connectWaitTimeSeconds = 3
	}

	return &gormDatabaseConnection{
		dsn:                    dsn,
		connectAttempts:        connectAttempts,
		connectWaitTimeSeconds: connectWaitTimeSeconds,
		connectBlocks:          connectBlocks,
		connectFinished:        make(chan bool, 1),
		status:                 db.UNKNOWN,
		logger:                 logger,
	}
}

func (c *gormDatabaseConnection) Connection() (*gorm.DB, error) {
	c.updateDBStatus()

	if c.status != db.READY {
		c.tryConnect()
	}

	if c.status == db.READY {
		return c.conn, nil
	}

	return nil, fmt.Errorf("database connection is not ready: %v, %v", c.status, c.connectError)
}

// tryConnect attempts opening single connection to the database.
func (c *gormDatabaseConnection) tryConnect() {
	if !c.connecting {
		if c.connectBlocks {
			c.connect()
		} else {
			go c.connect()
		}
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
		for c.status != db.READY {
			c.conn, c.connectError = gormConnect()
			c.updateDBStatus()

			if c.status != db.READY {
				c.logger.Warn(fmt.Sprintf("unable to connect to database: %v. retrying after %d seconds", c.connectError, c.connectWaitTimeSeconds))
				sleep(c.connectWaitTimeSeconds)
			}
		}

		c.logger.Info("connected with gorm to postgres")
	} else {
		var err error
		for i := 0; i < c.connectAttempts; i++ {
			c.conn, err = gormConnect()
			c.updateDBStatus()

			if c.status != db.READY {
				c.logger.Warn(fmt.Sprintf("unable to connect to database: %v", err))

				if i < c.connectAttempts-1 {
					sleep(c.connectWaitTimeSeconds)
				}
			} else {
				c.logger.Info("connected with gorm to postgres")
				break
			}
		}

		if c.isConnNil() {
			c.logger.Error(fmt.Sprintf("failed to connect to database in %d tries: %v", c.connectAttempts, err))
			c.connectError = err
		}
	}

	c.connecting = false
	go func() {
		c.connectFinished <- true
	}()
}

func (c *gormDatabaseConnection) isConnNil() bool {
	return db.IsInterfaceNil(c.conn)
}

func (c *gormDatabaseConnection) updateDBStatus() {
	if c.isConnNil() {
		c.status = db.NOT_READY
		return
	}

	sqlDB, err := c.conn.DB()
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to get generic SQL database: %v", err))
		c.status = db.ERROR

		if err = c.Close(context.Background()); err == nil {
			c.status = db.NOT_READY
		}
		return
	}

	if err := sqlDB.PingContext(context.Background()); err != nil {
		c.logger.Error(fmt.Sprintf("failed to ping database: %v", err))
		c.status = db.ERROR

		if err = c.Close(context.Background()); err == nil {
			c.status = db.NOT_READY
		}
	} else {
		c.status = db.READY
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
		c.logger.Error(fmt.Sprintf("failed to get generic SQL database: %v", err))
		c.status = db.ERROR
	}

	err = sqlDB.Close()
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to close connection: %v", err))
	}

	c.status = db.DISCONNECTED
	c.conn = nil
	return err
}
