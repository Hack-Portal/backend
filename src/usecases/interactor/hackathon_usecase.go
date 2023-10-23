package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/src/domain/params"
	"github.com/hackhack-Geek-vol6/backend/src/domain/request"
	"github.com/hackhack-Geek-vol6/backend/src/domain/response"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
)

type hackathonUsecase struct {
	store   transaction.Store
	l       logger.Logger
	timeout time.Duration
}

func NewHackathonUsercase(store transaction.Store, l logger.Logger) inputport.HackathonUsecase {
	return &hackathonUsecase{
		store:   store,
		l:       l,
		timeout: time.Duration(config.Config.Server.ContextTimeout),
	}
}

func (hu *hackathonUsecase) CreateHackathon(ctx context.Context, body request.CreateHackathon, image []byte) (result response.Hackathon, err error) {
	ctx, cancel := context.WithTimeout(ctx, hu.timeout)
	defer cancel()

	var imageURL string
	if image != nil {
		var err error
		_, imageURL, err = hu.store.UploadImage(ctx, image)
		if err != nil {
			return response.Hackathon{}, err
		}
	}

	hackathon, err := hu.store.CreateHackathonTx(ctx, params.CreateHackathon{
		Hackathon: repository.CreateHackathonsParams{
			Name: body.Name,
			Icon: sql.NullString{
				String: imageURL,
				Valid:  true,
			},
			Description: body.Description,
			Link:        body.Link,
			Expired:     body.Expired,
			StartDate:   body.StartDate,
			Term:        body.Term,
		},
		StatusTags: body.StatusTags,
	})
	if err != nil {
		return
	}

	statusTags, err := getHackathonTag(ctx, hu.store, hackathon.HackathonID)
	if err != nil {
		return
	}

	result = response.Hackathon{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTags:  statusTags,
	}
	return
}

func (hu *hackathonUsecase) ListHackathons(ctx context.Context, query request.ListHackathons) (result []response.ListHackathons, err error) {
	ctx, cancel := context.WithTimeout(ctx, hu.timeout)
	defer cancel()

	var expired time.Time
	if query.Expired {
		expired = time.Now().Add(-time.Hour * 24 * 30 * 6)
	} else {
		expired = time.Now()
	}

	hackathons, err := hu.store.ListHackathons(ctx, repository.ListHackathonsParams{
		Expired: expired,
		Limit:   query.PageSize,
		Offset:  (query.PageID - 1) * query.PageSize,
	})
	if err != nil {
		return
	}

	hackathonIDs := make([]int32, len(hackathons))
	for i, hackathon := range hackathons {
		hackathonIDs[i] = hackathon.HackathonID
	}

	statusTags, err := hu.store.ListHackathonStatusTagsByIDs(ctx, hackathonIDs)

	for _, hackathon := range hackathons {
		var tags []repository.StatusTag

		statusTags, err := hu.store.ListHackathonStatusTagsByID(ctx, hackathon.HackathonID)
		if err != nil {
			return nil, err
		}

		for _, statusTag := range statusTags {
			tag, err := hu.store.GetStatusTagsByTag(ctx, statusTag.StatusID)
			if err != nil {
				return nil, err
			}
			tags = append(tags, tag)
		}

		result = append(result, response.ListHackathons{
			HackathonID: hackathon.HackathonID,
			Icon:        hackathon.Icon.String,
			Name:        hackathon.Name,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			StatusTags:  tags,
		})
	}

	return
}

func (hu *hackathonUsecase) GetHackathon(ctx context.Context, id int32) (result response.Hackathon, err error) {
	ctx, cancel := context.WithTimeout(ctx, hu.timeout)
	defer cancel()

	hackathon, err := hu.store.GetHackathonByID(ctx, id)
	if err != nil {
		return response.Hackathon{}, err
	}

	statusTags, err := hu.store.ListHackathonStatusTagsByID(ctx, hackathon.HackathonID)
	if err != nil {
		return
	}

	var tags []repository.StatusTag
	for _, statusTag := range statusTags {
		tag, err := hu.store.GetStatusTagsByTag(ctx, statusTag.StatusID)
		if err != nil {
			return response.Hackathon{}, err
		}
		tags = append(tags, tag)
	}

	result = response.Hackathon{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
		StatusTags:  tags,
	}
	return
}

func getHackathonTag(ctx context.Context, store transaction.Store, id int32) (result []repository.StatusTag, err error) {
	tags, err := store.ListHackathonStatusTagsByID(ctx, id)
	if err != nil {
		return
	}

	for _, tag := range tags {
		statusTag, err := store.GetStatusTagsByTag(ctx, tag.StatusID)
		if err != nil {
			return nil, err
		}
		result = append(result, statusTag)
	}
	return
}
