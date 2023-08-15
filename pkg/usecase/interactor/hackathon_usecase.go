package usecase

import (
	"context"
	"database/sql"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type hackathonUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewHackathonUsercase(store transaction.Store, timeout time.Duration) inputport.HackathonUsecase {
	return &hackathonUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (hu *hackathonUsecase) CreateHackathon(ctx context.Context, body domain.CreateHackathonRequestBody, image []byte) (result domain.HackathonResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, hu.contextTimeout)
	defer cancel()

	var imageURL string
	if image != nil {
		var err error
		imageURL, err = hu.store.UploadImage(ctx, image)
		if err != nil {
			return domain.HackathonResponses{}, err
		}
	}

	hackathon, err := hu.store.CreateHackathonTx(ctx, domain.CreateHackathonParams{
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

	result = domain.HackathonResponses{
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

func (hu *hackathonUsecase) ListHackathons(ctx context.Context, query domain.ListHackathonsParams) (result []domain.ListHackathonsResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, hu.contextTimeout)
	defer cancel()

	var expired time.Time
	if query.Expired {
		expired = time.Now().Add(time.Hour * 24 * 30 * 6)
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

		result = append(result, domain.ListHackathonsResponses{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			StatusTags:  tags,
		})
	}

	return
}

func (hu *hackathonUsecase) GetHackathon(ctx context.Context, id int32) (result domain.HackathonResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, hu.contextTimeout)
	defer cancel()

	hackathon, err := hu.store.GetHackathonByID(ctx, id)
	if err != nil {
		return domain.HackathonResponses{}, err
	}

	statusTags, err := hu.store.ListHackathonStatusTagsByID(ctx, hackathon.HackathonID)
	if err != nil {
		return
	}

	var tags []repository.StatusTag
	for _, statusTag := range statusTags {
		tag, err := hu.store.GetStatusTagsByTag(ctx, statusTag.StatusID)
		if err != nil {
			return domain.HackathonResponses{}, err
		}
		tags = append(tags, tag)
	}

	result = domain.HackathonResponses{
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
