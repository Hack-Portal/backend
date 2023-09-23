package inputport

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
)

type HackathonUsecase interface {
	CreateHackathon(ctx context.Context, body request.CreateHackathonRequestBody, image []byte) (result response.HackathonResponses, err error)
	GetHackathon(ctx context.Context, id int32) (result response.HackathonResponses, err error)
	ListHackathons(ctx context.Context, query request.ListHackathonsRequest) (result []response.ListHackathonsResponses, err error)
}
