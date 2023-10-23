package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/src/transaction"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
	usecase "github.com/hackhack-Geek-vol6/backend/src/usecases/interactor"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

type EtcController struct {
	EtcUsecase inputport.EtcUsecase
	l          logger.Logger
}

func NewEtcController(store transaction.SQLStore, l logger.Logger) *EtcController {
	return &EtcController{
		EtcUsecase: usecase.NewEtcUsercase(store, l),
		l:          l,
	}
}

// Health	godoc
//
//	@Summary		Health Check
//	@Description	Health Check
//	@Tags			Health
//	@Produce		json
//	@Success		200			{object}	HealthResponse	"success response"
//	@Failure		500			{object}	ErrorResponse		"error response"
//	@Router			/health		[get]
func (ec *EtcController) Health(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// ListFrameworks	godoc
//
//	@Summary		Get Frameworks
//	@Description	Get Frameworks
//	@Tags			Frameworks
//	@Produce		json
//	@Success		200			{array}		repository.Framework	"success response"
//	@Failure		500			{object}	ErrorResponse			"error response"
//	@Router			/frameworks	[get]
func (ec *EtcController) ListFrameworks(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	response, err := ec.EtcUsecase.GetFramework(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// ListLocation	godoc
//
//	@Summary		Get Frameworks
//	@Description	Get Frameworks
//	@Tags			Locates
//	@Produce		json
//	@Success		200			{array}		repository.Locate	"success response"
//	@Failure		500			{object}	ErrorResponse		"error response"
//	@Router			/locates	[get]
func (ec *EtcController) ListLocation(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	response, err := ec.EtcUsecase.GetLocat(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// ListTechTags		godoc
//
//	@Summary		Get Frameworks
//	@Description	Get Frameworks
//	@Tags			TechTags
//	@Produce		json
//	@Success		200			{array}		repository.TechTag	"success response"
//	@Failure		500			{object}	ErrorResponse		"error response"
//	@Router			/tech_tags	[get]
func (ec *EtcController) ListTechTags(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	response, err := ec.EtcUsecase.GetTechTag(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// ListStatusTags	godoc
//
//	@Summary		Get Frameworks
//	@Description	Get Frameworks
//	@Tags			TechTags
//	@Produce		json
//	@Success		200				{array}		repository.StatusTag	"success response"
//	@Failure		500				{object}	ErrorResponse			"error response"
//	@Router			/status_tags	[get]
func (ec *EtcController) ListStatusTags(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	response, err := ec.EtcUsecase.GetStatusTag(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (ec *EtcController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// ListRoles	godoc
//
//	@Summary		Get Roles
//	@Description	Get Roles
//	@Tags			Roles
//	@Produce		json
//	@Success		200				{array}		repository.Role	"success response"
//	@Failure		500				{object}	ErrorResponse		"error response"
//	@Router			/roles																										[get]
func (ec *EtcController) ListRoles(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	response, err := ec.EtcUsecase.ListRoles(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
