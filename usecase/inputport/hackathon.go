package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/domain"
)

type HackathonUsecase interface {
	CreateHackathon(ctx context.Context, body domain.CreateHackathonParams) (result domain.HackathonResponses, err error)
	GetHackathon(ctx context.Context, id int32) (result domain.HackathonResponses, err error)
	ListHackathons(ctx context.Context, query domain.ListHackathonsParams) (result []domain.ListHackathonsResponses, err error)
}
