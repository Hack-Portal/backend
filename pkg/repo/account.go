package repo

import (
	"database/sql"

	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
)

type AccountRepository interface {
}

type accountRepository struct {
	db *sql.DB
	l  logger.Logger
}

func NewAccountRepository(db *sql.DB, l logger.Logger) AccountRepository {
	return &accountRepository{
		db: db,
		l:  l,
	}
}
