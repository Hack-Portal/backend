package transaction

import (
	"context"
	"fmt"

	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
)

func (store *SQLStore) execTx(ctx context.Context, fn func(*repository.Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := repository.New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}
