package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/usecase/inputport"
)

type etcUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewEtcUsercase(store transaction.Store, timeout time.Duration) inputport.EtcUsecase {
	return &etcUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (eu *etcUsecase) GetFramework(ctx context.Context, limit int32) ([]repository.Framework, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListFrameworks(ctx, limit)
}

func (eu *etcUsecase) GetLocat(ctx context.Context) ([]repository.Locate, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListLocates(ctx)
}
func (eu *etcUsecase) GetTechTag(ctx context.Context) ([]repository.TechTag, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListTechTag(ctx)
}
