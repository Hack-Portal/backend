package transaction

import (
	"context"
	"database/sql"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/datasource"
)

func parseAccountResponse(response domain.AccountResponses, account repository.Account, locate string) domain.AccountResponses {
	return domain.AccountResponses{
		UserID:          account.UserID,
		Username:        account.Username,
		Icon:            account.Icon.String,
		ExplanatoryText: account.ExplanatoryText.String,
		Rate:            account.Rate,
		Email:           account.Email,
		Locate:          locate,
		ShowRate:        account.ShowRate,
		ShowLocate:      account.ShowLocate,
		TechTags:        response.TechTags,
		Frameworks:      response.Frameworks,
	}
}

func createTagsAndFrameworks(ctx context.Context, q *repository.Queries, id string, techTags, frameworks []int32) (tag []repository.TechTag, fw []repository.Framework, err error) {
	for _, techTag := range techTags {
		accountTag, err := q.CreateAccountTags(ctx, repository.CreateAccountTagsParams{
			UserID:    id,
			TechTagID: techTag,
		})
		if err != nil {
			return nil, nil, err
		}

		techtag, err := q.GetTechTagByID(ctx, accountTag.TechTagID)
		if err != nil {
			return nil, nil, err
		}

		tag = append(tag, techtag)
	}

	for _, accountFrameworkTag := range frameworks {
		accountFramework, err := q.CreateAccountFramework(ctx, repository.CreateAccountFrameworkParams{
			UserID:      id,
			FrameworkID: accountFrameworkTag,
		})
		if err != nil {
			return nil, nil, err
		}

		framework, err := q.GetFrameworksByID(ctx, accountFramework.FrameworkID)
		if err != nil {
			return nil, nil, err
		}
		fw = append(fw, framework)
	}
	return
}

func parseAccountResponseRawID(response domain.AccountResponses, account repository.GetAccountByIDRow, locate string) domain.AccountResponses {
	return domain.AccountResponses{
		UserID:          account.UserID,
		Username:        account.Username,
		Icon:            account.Icon.String,
		ExplanatoryText: account.ExplanatoryText.String,
		Rate:            account.Rate,
		Email:           account.Email,
		Locate:          locate,
		ShowRate:        account.ShowRate,
		ShowLocate:      account.ShowLocate,
		TechTags:        response.TechTags,
		Frameworks:      response.Frameworks,
	}
}

func parseAccountResponseRawEmail(response domain.AccountResponses, account repository.GetAccountByEmailRow, locate string) domain.AccountResponses {
	return domain.AccountResponses{
		UserID:          account.UserID,
		Username:        account.Username,
		Icon:            account.Icon.String,
		ExplanatoryText: account.ExplanatoryText.String,
		Rate:            account.Rate,
		Email:           account.Email,
		Locate:          locate,
		ShowRate:        account.ShowRate,
		ShowLocate:      account.ShowLocate,
		TechTags:        response.TechTags,
		Frameworks:      response.Frameworks,
	}
}

func compAccount(request repository.Account, latest repository.GetAccountByIDRow) (result repository.UpdateAccountParams) {
	result = repository.UpdateAccountParams{
		UserID:         latest.UserID,
		Icon:           latest.Icon,
		Rate:           latest.Rate,
		HashedPassword: latest.HashedPassword,
		Email:          latest.Email,
	}

	if len(request.Username) != 0 {
		if latest.Username != request.Username {
			result.Username = request.Username
		}
	} else {
		result.Username = latest.Username
	}

	if len(request.ExplanatoryText.String) != 0 {
		if latest.ExplanatoryText.String != request.ExplanatoryText.String {
			result.ExplanatoryText = sql.NullString{
				String: request.ExplanatoryText.String,
				Valid:  true,
			}
		}
	} else {
		result.ExplanatoryText = latest.ExplanatoryText
	}

	if len(request.Icon.String) != 0 {
		if latest.Icon.String != request.Icon.String {
			result.Icon = sql.NullString{
				String: request.Icon.String,
				Valid:  true,
			}
		}
	} else {
		result.Icon = latest.Icon
	}

	if request.LocateID != 0 {
		if latest.LocateID != request.LocateID {
			result.LocateID = request.LocateID
		}
	} else {
		result.LocateID = latest.LocateID
	}

	if latest.ShowLocate != request.ShowLocate {
		result.ShowLocate = request.ShowLocate
	} else {
		result.ShowLocate = latest.ShowLocate
	}

	if latest.ShowRate != request.ShowRate {
		result.ShowRate = request.ShowRate
	} else {
		result.ShowRate = latest.ShowRate
	}

	return
}

func (store *SQLStore) CreateAccountTx(ctx context.Context, args domain.CreateAccountParams) (domain.AccountResponses, error) {
	var result domain.AccountResponses
	err := store.execTx(ctx, func(q *repository.Queries) error {

		account, err := q.CreateAccount(ctx, args.AccountInfo)
		if err != nil {
			return err
		}

		locate, err := q.GetLocateByID(ctx, account.LocateID)
		if err != nil {
			return err
		}

		techTags, frameworks, err := createTagsAndFrameworks(ctx, q, args.AccountInfo.UserID, args.AccountTechTag, args.AccountFrameworkTag)
		if err != nil {
			return err
		}

		result = parseAccountResponse(result, account, locate.Name)
		result.TechTags = techTags
		result.Frameworks = frameworks

		return nil
	})
	return result, err
}

func getAccountTags(ctx context.Context, q *repository.Queries, id string) (rTechTags []repository.TechTag, rFrameworks []repository.Framework, err error) {
	techTags, err := q.ListAccountTagsByUserID(ctx, id)
	if err != nil {
		return
	}

	frameworks, err := q.ListAccountFrameworksByUserID(ctx, id)
	if err != nil {
		return
	}

	for _, techTag := range techTags {
		techtag, err := q.GetTechTagByID(ctx, techTag.TechTagID.Int32)
		if err != nil {
			return nil, nil, err
		}

		rTechTags = append(rTechTags, techtag)
	}

	for _, framework := range frameworks {
		framework, err := q.GetFrameworksByID(ctx, framework.FrameworkID.Int32)
		if err != nil {
			return nil, nil, err
		}
		rFrameworks = append(rFrameworks, framework)
	}

	return
}

func (store *SQLStore) GetAccountTxByID(ctx context.Context, id string) (domain.AccountResponses, error) {
	var result domain.AccountResponses
	err := store.execTx(ctx, func(q *repository.Queries) error {

		account, err := q.GetAccountByID(ctx, id)
		if err != nil {
			return err
		}

		locate, err := q.GetLocateByID(ctx, account.LocateID)
		if err != nil {
			return err
		}

		techTags, frameworks, err := getAccountTags(ctx, q, account.UserID)

		result = parseAccountResponseRawID(result, account, locate.Name)
		result.TechTags = techTags
		result.Frameworks = frameworks

		return nil
	})
	return result, err
}

func (store *SQLStore) GetAccountTxByEmail(ctx context.Context, email string) (domain.AccountResponses, error) {
	var result domain.AccountResponses
	err := store.execTx(ctx, func(q *repository.Queries) error {

		account, err := q.GetAccountByEmail(ctx, email)
		if err != nil {
			return err
		}

		locate, err := q.GetLocateByID(ctx, account.LocateID)
		if err != nil {
			return err
		}

		techTags, frameworks, err := getAccountTags(ctx, q, account.UserID)
		result = parseAccountResponseRawEmail(result, account, locate.Name)
		result.TechTags = techTags
		result.Frameworks = frameworks

		return nil
	})
	return result, err
}

// アカウントの新旧の比較をする

func (store *SQLStore) UpdateAccountTx(ctx context.Context, args domain.UpdateAccountParam) (domain.AccountResponses, error) {
	var result domain.AccountResponses
	err := store.execTx(ctx, func(q *repository.Queries) error {

		latest, err := q.GetAccountByID(ctx, args.AccountInfo.UserID)
		if err != nil {
			return err
		}

		account, err := q.UpdateAccount(ctx, compAccount(args.AccountInfo, latest))
		if err != nil {
			return err
		}

		locate, err := q.GetLocateByID(ctx, account.LocateID)
		if err != nil {
			return err
		}

		// 以下タグ部分
		err = q.DeleteAccounttagsByUserID(ctx, latest.UserID)
		if err != nil {
			return err
		}

		err = q.DeleteAccountFrameworksByUserID(ctx, latest.UserID)
		if err != nil {
			return err
		}

		techTags, frameworks, err := createTagsAndFrameworks(ctx, q, latest.UserID, args.AccountTechTag, args.AccountFrameworkTag)
		if err != nil {
			return err
		}

		result = parseAccountResponse(result, account, locate.Name)
		result.TechTags = techTags
		result.Frameworks = frameworks

		return nil
	})
	return result, err
}