package transaction

import (
	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	repository.Querier
	// Transactios interface
}

type SQLStore struct {
	connPool *pgxpool.Pool
	*repository.Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  repository.New(connPool),
	}
}

func BeginTx()
