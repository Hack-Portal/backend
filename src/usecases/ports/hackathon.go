package ports

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
)

type HackathonInputBoundary interface {
	CreateHackathon(ctx context.Context, in *InputCreatehackathonData) (int, *response.CreateHackathon)
	GetHackathon(ctx context.Context, hackathonID string) (int, *response.GetHackathon)
	ListHackathon(ctx context.Context, in request.ListHackathon) (int, []*response.GetHackathon)
	DeleteHackathon(ctx context.Context, hackathonID string) (int, *response.DeleteHackathon)
}

type HackathonOutputBoundary interface {
	PresentCreateHackathon(ctx context.Context, out *OutputCreateHackathonData) (int, *response.CreateHackathon)
	PresentGetHackathon(ctx context.Context, out *OutputGetHackathonData) (int, *response.GetHackathon)
	PresentListHackathon(ctx context.Context, out *OutputListHackathonData) (int, []*response.GetHackathon)
	PresentDeleteHackathon(ctx context.Context, out *OutputDeleteHackathonData) (int, *response.DeleteHackathon)
}

type InputCreatehackathonData struct {
	ImageFile *multipart.FileHeader

	Name      string
	Link      string
	Expired   time.Time
	StartDate time.Time
	Term      int
	Statuses  []int64
}

type OutputCreateHackathonData struct {
	Error    error
	Response *response.CreateHackathon
}

type OutputGetHackathonData struct {
	Error    error
	Response *response.GetHackathon
}

type OutputListHackathonData struct {
	Error    error
	Response []*response.GetHackathon
}

type OutputDeleteHackathonData struct {
	Error    error
	Response *response.DeleteHackathon
}
