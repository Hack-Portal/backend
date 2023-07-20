package api

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

func (server *Server) Ping(ctx *gin.Context) {
	var imageURL string
	file, _, err := ctx.Request.FormFile(ImageKey)
	if err != nil {
		switch err.Error() {
		case MultiPartNextPartEoF:
			ctx.JSON(400, errorResponse(err))
			return
		case HttpNoSuchFile:
			ctx.JSON(400, errorResponse(err))
			return
		default:
			ctx.JSON(400, errorResponse(err))
			return
		case RequestContentTypeIsnt:
			break
		}
	} else {
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			ctx.JSON(500, errorResponse(err))
			return
		}
		imageURL, err = server.store.UploadImage(ctx, buf.Bytes())
		if _, err := io.Copy(buf, file); err != nil {
			ctx.JSON(500, errorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": imageURL})
}

func (server *Server) Pong(ctx *gin.Context) {
	authClaims := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	ctx.JSON(http.StatusOK, gin.H{"message": "ping", "claims": authClaims})
}
