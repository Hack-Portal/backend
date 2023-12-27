package ports

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/Hack-Portal/backend/src/datastructure/response"
)

type HackathonInputBoundary interface {
	CreateHackathon(ctx context.Context, in *InputCreatehackathonData) (int, *response.CreateHackathon)
}

type HackathonOutputBoundary interface {
	PresentCreateHackathon(ctx context.Context, out *OutputCreateHackathonData) (int, *response.CreateHackathon)
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
