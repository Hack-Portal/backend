package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Store interface {
	Querier
	// Transactios interface
}

type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
