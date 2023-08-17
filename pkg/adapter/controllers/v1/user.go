package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type UserController struct {
	UserUsecase inputport.UserUsecase
	Env         *bootstrap.Env
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var (
		reqBody domain.CreateUserRequest
	)
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	duration, err := strconv.Atoi(uc.Env.TokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response, err := uc.UserUsecase.CreateUser(ctx, reqBody, time.Duration(duration)*time.Hour*24)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
