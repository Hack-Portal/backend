package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type pastWorkUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewPastWorkUsercase(store transaction.Store, timeout time.Duration) inputport.PastworkUsecase {
	return &pastWorkUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (pu *pastWorkUsecase) CreatePastWork(ctx context.Context, arg params.CreatePastWork, image []byte) (result response.PastWorkResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	// 画像が空でないときに処理する
	var imageURL string
	if image != nil {
		_, imageURL, err = pu.store.UploadImage(ctx, image)
		if err != nil {
			return response.PastWorkResponse{}, err
		}
	}
	pastWork, err := pu.store.CreatePastWorkTx(ctx, params.CreatePastWork{
		Name:               arg.Name,
		ThumbnailImage:     imageURL,
		ExplanatoryText:    arg.ExplanatoryText,
		PastWorkTags:       arg.PastWorkTags,
		PastWorkFrameworks: arg.PastWorkFrameworks,
		AccountPastWorks:   arg.AccountPastWorks,
	})
	if err != nil {
		return response.PastWorkResponse{}, err
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

	result = response.PastWorkResponse{
		Opus:            pastWork.Opus,
		Name:            pastWork.Name,
		ThumbnailImage:  pastWork.ThumbnailImage,
		ExplanatoryText: pastWork.ExplanatoryText,
		AwardDataID:     pastWork.AwardDataID.Int32,
		CreateAt:        pastWork.CreateAt,
		UpdateAt:        pastWork.UpdateAt,
		TechTags:        techTags,
		Frameworks:      frameworks,
		Members:         members,
	}
	return
}

func (pu *pastWorkUsecase) GetPastWork(ctx context.Context, opus int32) (result response.PastWorkResponse, err error) {
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

	result = response.PastWorkResponse{
		Opus:            pastWork.Opus,
		Name:            pastWork.Name,
		ThumbnailImage:  pastWork.ThumbnailImage,
		ExplanatoryText: pastWork.ExplanatoryText,
		AwardDataID:     pastWork.AwardDataID.Int32,
		CreateAt:        pastWork.CreateAt,
		UpdateAt:        pastWork.UpdateAt,
		TechTags:        techTags,
		Frameworks:      frameworks,
		Members:         members,
	}
	return
}

func (pu *pastWorkUsecase) ListPastWork(ctx context.Context, query request.ListRequest) (result []response.ListPastWorkResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	pastWorks, err := pu.store.ListPastWorks(ctx, repository.ListPastWorksParams{Limit: query.PageSize, Offset: (query.PageID - 1) * query.PageSize})
	if err != nil {
		return
	}

	for _, pastWork := range pastWorks {
		techTags, err := parsePastWorkTechTags(ctx, pu.store, pastWork.Opus)
		if err != nil {
			return nil, err
		}

		frameworks, err := parsePastWorkFrameworks(ctx, pu.store, pastWork.Opus)
		if err != nil {
			return nil, err
		}

		members, err := parsePastWorkMembers(ctx, pu.store, pastWork.Opus)
		if err != nil {
			return nil, err
		}

		result = append(result, response.ListPastWorkResponse{
			Opus:            pastWork.Opus,
			Name:            pastWork.Name,
			ExplanatoryText: pastWork.ExplanatoryText,
			TechTags:        techTags,
			Frameworks:      frameworks,
			Members:         members,
		})
	}
	return
}

func (pu *pastWorkUsecase) UpdatePastWork(ctx context.Context, body params.UpdatePastWork) (result response.PastWorkResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	pastWork, err := pu.store.UpdatePastWorkTx(ctx, body)
	if err != nil {
		return
	}

	techTags, err := parsePastWorkTechTags(ctx, pu.store, pastWork.Opus)
	if err != nil {
		return
	}

	frameworks, err := parsePastWorkFrameworks(ctx, pu.store, pastWork.Opus)
	if err != nil {
		return response.PastWorkResponse{}, err
	}

	members, err := parsePastWorkMembers(ctx, pu.store, pastWork.Opus)
	if err != nil {
		return response.PastWorkResponse{}, err
	}

	result = response.PastWorkResponse{
		Opus:            pastWork.Opus,
		Name:            pastWork.Name,
		ThumbnailImage:  pastWork.ThumbnailImage,
		ExplanatoryText: pastWork.ExplanatoryText,
		AwardDataID:     pastWork.AwardDataID.Int32,
		CreateAt:        pastWork.CreateAt,
		UpdateAt:        pastWork.UpdateAt,
		TechTags:        techTags,
		Frameworks:      frameworks,
		Members:         members,
	}
	return result, nil
}

func (pu *pastWorkUsecase) DeletePastWork(ctx context.Context, args repository.DeletePastWorksByIDParams) error {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	_, err := pu.store.DeletePastWorksByID(ctx, args)
	return err
}

func parsePastWorkMembers(ctx context.Context, store transaction.Store, opus int32) (result []response.PastWorkMembers, err error) {
	members, err := store.ListAccountPastWorksByOpus(ctx, opus)
	if err != nil {
		return
	}

	for _, member := range members {
		account, err := store.GetAccountsByID(ctx, member.AccountID)
		if err != nil {
			return nil, err
		}
		result = append(result, response.PastWorkMembers{AccountID: account.AccountID, Name: account.Username, Icon: account.Icon.String})
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
