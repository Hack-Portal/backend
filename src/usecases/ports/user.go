package ports

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
)

type UserInputBoundary interface {
	InitAdmin(ctx context.Context, in request.InitAdmin) (int, *response.User)
	Login(ctx context.Context, in request.Login) (int, *response.Login)
}

type UserOutputBoundary interface {
	PresentInitAdmin(ctx context.Context, out *OutputInitAdminData) (int, *response.User)
	PresentLogin(ctx context.Context, out OutputBoundary[*response.Login]) (int, *response.Login)
}

type OutputInitAdminData struct {
	Error    error
	Response *response.User
}
