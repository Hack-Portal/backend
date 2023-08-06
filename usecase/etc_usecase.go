package usecase

import (
	"context"
	"time"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type etcUsecase struct {
	store          db.Store
	contextTimeout time.Duration
}

func NewEtcUsercase(store db.Store, timeout time.Duration) domain.EtcUsecase {
	return &etcUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (eu *etcUsecase) GetFramework(ctx context.Context, limit int32) ([]db.Frameworks, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListFrameworks(ctx, limit)
}

func (eu *etcUsecase) GetLocat(ctx context.Context) ([]db.Locates, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListLocates(ctx)
}
func (eu *etcUsecase) GetTechTag(ctx context.Context) ([]db.TechTags, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()

	return eu.store.ListTechTag(ctx)
}
