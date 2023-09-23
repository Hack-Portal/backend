package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
)

type HackathonUsecase interface {
	CreateHackathon(ctx context.Context, body request.CreateHackathon, image []byte) (result response.Hackathon, err error)
	GetHackathon(ctx context.Context, id int32) (result response.Hackathon, err error)
	ListHackathons(ctx context.Context, query request.ListHackathons) (result []response.ListHackathons, err error)
}
