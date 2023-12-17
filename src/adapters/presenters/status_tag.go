package presenters

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/hackhack-Geek-vol6/backend/src/datastructure/hperror"
	"github.com/hackhack-Geek-vol6/backend/src/datastructure/response"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/ports"
	"gorm.io/gorm"
)

type StatusTagPresenter struct {
	logger *slog.Logger
}

func NewStatusTagPresenter() ports.StatusTagOutputBoundary {
	return &StatusTagPresenter{}
}

func (s *StatusTagPresenter) PresentCreateStatusTag(ctx context.Context, out *ports.OutputCraeteStatusTagData) (int, *response.StatusTag) {
	if out.Error != nil {
		switch out.Error {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusCreated, &response.StatusTag{
		ID:     out.Response.StatusID,
		Status: out.Response.Status,
	}
}

func (s *StatusTagPresenter) PresentFindAllStatusTag(ctx context.Context, out *ports.OutputFindAllStatusTagData) (int, []*response.StatusTag) {
	if out.Error != nil {
		switch out.Error {
		case gorm.ErrRecordNotFound:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	var resp []*response.StatusTag
	for _, statusTag := range out.Response {
		resp = append(resp, &response.StatusTag{
			ID:     statusTag.StatusID,
			Status: statusTag.Status,
		})
	}

	return http.StatusOK, resp
}

func (s *StatusTagPresenter) PresentFindByIdStatusTag(ctx context.Context, out *ports.OutputFindByIdStatusTagData) (int, *response.StatusTag) {
	if out.Error != nil {
		switch out.Error {
		case gorm.ErrRecordNotFound:
			return http.StatusBadRequest, nil
		default:

			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusOK, &response.StatusTag{
		ID:     out.Response.StatusID,
		Status: out.Response.Status,
	}
}

func (s *StatusTagPresenter) PresentUpdateStatusTag(ctx context.Context, out *ports.OutputUpdateStatusTagData) (int, *response.StatusTag) {
	if out.Error != nil {
		switch out.Error {
		case gorm.ErrRecordNotFound:
			return http.StatusBadRequest, nil
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			// TODO: ここにエラーの種類を追加していく
			return http.StatusInternalServerError, nil
		}
	}

	return http.StatusOK, &response.StatusTag{
		ID:     out.Response.StatusID,
		Status: out.Response.Status,
	}
}
