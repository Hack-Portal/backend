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

// ListRooms	godoc
// @Summary			Create User
// @Description		Create User
// @Tags			Users
// @Produce			json
// @Param			CreateUserRequest 	body 	domain.CreateUserRequest	true				"Create User Request"
// @Success			200		{object}	domain.CreateUserResponse								"success response"
// @Failure 		400		{object}	ErrorResponse											"error response"
// @Failure 		500		{object}	ErrorResponse											"error response"
// @Router       	/users	[post]
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

// ListRooms		godoc
// @Summary			Login User
// @Description		Login User
// @Tags			Users
// @Produce			json
// @Param			CreateUserRequest 	body 		domain.CreateUserRequest	true	"List Rooms Request"
// @Success			200		{object}		domain.CreateUserResponse	"success response"
// @Failure 		400		{object}	ErrorResponse			"error response"
// @Failure 		500		{object}	ErrorResponse			"error response"
// @Router       	/login	[post]
func (uc *UserController) LoginUser(ctx *gin.Context) {
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

	response, err := uc.UserUsecase.LoginUser(ctx, reqBody, time.Duration(duration)*time.Hour*24)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
