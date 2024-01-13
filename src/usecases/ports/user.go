package ports

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
)

// UserInputBoundary はUserのInputBoundary
type UserInputBoundary interface {
	InitAdmin(ctx context.Context, in request.InitAdmin) (int, *response.User)
	Login(ctx context.Context, in request.Login) (int, *response.Login)
}

// UserOutputBoundary はUserのOutputBoundary
type UserOutputBoundary interface {
	PresentInitAdmin(ctx context.Context, out OutputBoundary[*response.User]) (int, *response.User)
	PresentLogin(ctx context.Context, out OutputBoundary[*response.Login]) (int, *response.Login)
}

// OutputInitAdminData はInitAdminのOutputData
type OutputInitAdminData struct {
	Error    error
	Response *response.User
}
