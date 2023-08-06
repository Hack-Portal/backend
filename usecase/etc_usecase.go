package usecase

import (
	"context"
	"time"

	"github.com/hackhack-Geek-vol6/backend/domain"
	"github.com/hackhack-Geek-vol6/backend/gateways/repository"
)

type etcUsecase struct {
	store          repository.Store
	contextTimeout time.Duration
}

func NewEtcUsercase(store repository.Store, timeout time.Duration) domain.EtcUsecase {
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
