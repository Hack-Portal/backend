package ports

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
)

// HackathonInputBoundary はHackathonのInputBoundary
type HackathonInputBoundary interface {
	CreateHackathon(ctx context.Context, in *InputCreatehackathonData) (int, *response.CreateHackathon)
	ListHackathon(ctx context.Context, in request.ListHackathon) (int, []*response.GetHackathon)
	DeleteHackathon(ctx context.Context, hackathonID string) (int, *response.DeleteHackathon)
}

// HackathonOutputBoundary はHackathonのOutputBoundary
type HackathonOutputBoundary interface {
	PresentCreateHackathon(ctx context.Context, out OutputBoundary[*response.CreateHackathon]) (int, *response.CreateHackathon)
	PresentListHackathon(ctx context.Context, out OutputBoundary[[]*response.GetHackathon]) (int, []*response.GetHackathon)
	PresentDeleteHackathon(ctx context.Context, out OutputBoundary[*response.DeleteHackathon]) (int, *response.DeleteHackathon)
}

// InputCreatehackathonData はHackathonのCreateHackathonのInputData
type InputCreatehackathonData struct {
	ImageFile *multipart.FileHeader

	Name      string
	Link      string
	Expired   time.Time
	StartDate time.Time
	Term      int
	Statuses  []int64
}
