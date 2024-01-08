package ports

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
)

type UserInputBoundary interface {
	InitAdmin(ctx context.Context, in request.InitAdmin) (int, *response.User)
}

type UserOutputBoundary interface {
	PresentInitAdmin(ctx context.Context, out *OutputInitAdminData) (int, *response.User)
}

type OutputInitAdminData struct {
	Error    error
	Response *response.User
}
