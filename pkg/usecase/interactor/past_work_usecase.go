package usecase

import (
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	"golang.org/x/net/context"
)

type pastWorkUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewPastWorkUsercase(store transaction.Store, timeout time.Duration) inputport.PastworksUsecase {
	return &pastWorkUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (pu *pastWorkUsecase) CreatePastWork(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

}

func (pu *pastWorkUsecase) GetPastWork(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

}

func (pu *pastWorkUsecase) ListPastWork(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

}

func (pu *pastWorkUsecase) UpdatePastWork(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

}

func (pu *pastWorkUsecase) DeletePastWork(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

}
