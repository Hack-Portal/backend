package usecase

import (
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
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

func (pu *pastWorkUsecase) CreatePastWork(ctx context.Context, arg domain.CreatePastWorkParams) (result domain.PastWorkResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	pastWork, err := pu.store.CreatePastWorkTx(ctx, arg)
	if err != nil {
		return
	}

	techTags, err := parsePastWorkTechTags(ctx, pu.store, pastWork.Opus)
	if err != nil {
		return
	}

	frameworks, err := parsePastWorkFrameworks(ctx, pu.store, pastWork.Opus)
	if err != nil {
		return
	}

	members, err := parsePastWorkMembers(ctx, pu.store, pastWork.Opus)
	if err != nil {
		return
	}

	result = domain.PastWorkResponse{
		Pastwork:   pastWork,
		TechTags:   techTags,
		Frameworks: frameworks,
		Members:    members,
	}
	return
}

func (pu *pastWorkUsecase) GetPastWork(ctx context.Context, opus int32) (result domain.PastWorkResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	pastWork, err := pu.store.GetPastWorksByOpus(ctx, opus)
	if err != nil {
		return
	}

	techTags, err := parsePastWorkTechTags(ctx, pu.store, pastWork.Opus)
	if err != nil {
		return
	}

	frameworks, err := parsePastWorkFrameworks(ctx, pu.store, pastWork.Opus)
	if err != nil {
		return
	}

	members, err := parsePastWorkMembers(ctx, pu.store, pastWork.Opus)
	if err != nil {
		return
	}

	result = domain.PastWorkResponse{
		Pastwork:   pastWork,
		TechTags:   techTags,
		Frameworks: frameworks,
		Members:    members,
	}
	return
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

func parsePastWorkMembers(ctx context.Context, store transaction.Store, opus int32) (result []domain.PastWorkMembers, err error) {
	members, err := store.ListAccountPastWorksByOpus(ctx, opus)
	if err != nil {
		return
	}

	for _, member := range members {
		account, err := store.GetAccountsByID(ctx, member.AccountID)
		if err != nil {
			return nil, err
		}
		result = append(result, domain.PastWorkMembers{AccountID: account.AccountID, Name: account.Username, Icon: account.Icon.String})
	}
	return
}

func parsePastWorkTechTags(ctx context.Context, store transaction.Store, opus int32) (result []repository.TechTag, err error) {
	techTags, err := store.ListPastWorkTagsByOpus(ctx, opus)
	if err != nil {
		return
	}

	for _, techTag := range techTags {
		tag, err := store.GetTechTagsByID(ctx, techTag.TechTagID)
		if err != nil {
			return nil, err
		}
		result = append(result, tag)
	}
	return
}

func parsePastWorkFrameworks(ctx context.Context, store transaction.Store, opus int32) (result []repository.Framework, err error) {
	frameworks, err := store.ListPastWorkFrameworksByOpus(ctx, opus)
	if err != nil {
		return
	}
	for _, framework := range frameworks {
		fw, err := store.GetFrameworksByID(ctx, framework.FrameworkID)
		if err != nil {
			return nil, err
		}
		result = append(result, fw)
	}
	return
}
