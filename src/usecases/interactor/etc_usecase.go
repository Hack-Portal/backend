package usecase

import (
	"context"
	"time"

	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
)

type etcUsecase struct {
	store   transaction.Store
	l       logger.Logger
	timeout time.Duration
}

func NewEtcUsercase(store transaction.Store, l logger.Logger) inputport.EtcUsecase {
	return &etcUsecase{
		store:   store,
		l:       l,
		timeout: time.Duration(config.Config.Server.ContextTimeout),
	}
}

func (eu *etcUsecase) GetFramework(ctx context.Context) ([]repository.Framework, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.timeout)
	defer cancel()

	return eu.store.ListFrameworks(ctx)
}

func (eu *etcUsecase) GetLocat(ctx context.Context) ([]repository.Locate, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.timeout)
	defer cancel()

	return eu.store.ListLocates(ctx)
}

func (eu *etcUsecase) GetTechTag(ctx context.Context) ([]repository.TechTag, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.timeout)
	defer cancel()

	return eu.store.ListTechTags(ctx)
}

func (eu *etcUsecase) GetStatusTag(ctx context.Context) ([]repository.StatusTag, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.timeout)
	defer cancel()

	return eu.store.ListStatusTags(ctx)
}

func (eu *etcUsecase) ListRoles(ctx context.Context) ([]repository.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.timeout)
	defer cancel()

	return eu.store.ListRoles(ctx)
}
