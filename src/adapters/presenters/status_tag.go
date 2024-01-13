package presenters

import (
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type statusTagPresenter struct{}

// NewStatusTagPresenter はStatusTagPresenterを返す
func NewStatusTagPresenter() ports.StatusTagOutputBoundary {
	return &statusTagPresenter{}
}

// PresentCreateStatusTag はStatusTagの作成をpresenterする
func (s *statusTagPresenter) PresentCreateStatusTag(ctx context.Context, out ports.OutputBoundary[*models.StatusTag]) (int, *response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("PresentCreateStatusTag-presenter").End()

	if err := out.Error(); err != nil {
		switch out.Error() {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	resp := out.Response()

	return http.StatusCreated, &response.StatusTag{
		ID:     resp.StatusID,
		Status: resp.Status,
	}
}

// PresentFindAllStatusTag はStatusTagの取得をpresenterする
func (s *statusTagPresenter) PresentFindAllStatusTag(ctx context.Context, out ports.OutputBoundary[[]*models.StatusTag]) (int, []*response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("PresentFindAllStatusTag-presenter").End()

	if err := out.Error(); err != nil {
		switch out.Error() {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}
	var resp []*response.StatusTag
	for _, statusTag := range out.Response() {
		resp = append(resp, &response.StatusTag{
			ID:     statusTag.StatusID,
			Status: statusTag.Status,
		})
	}

	return http.StatusOK, resp
}

// PresentFindByIDStatusTag はStatusTagの取得をpresenterする
func (s *statusTagPresenter) PresentFindByIDStatusTag(ctx context.Context, out ports.OutputBoundary[*models.StatusTag]) (int, *response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("PresentFindByIdStatusTag-presenter").End()

	if err := out.Error(); err != nil {
		switch out.Error() {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	resp := out.Response()

	return http.StatusOK, &response.StatusTag{
		ID:     resp.StatusID,
		Status: resp.Status,
	}
}

// PresentDeleteStatusTag はStatusTagの削除をpresenterする
func (s *statusTagPresenter) PresentUpdateStatusTag(ctx context.Context, out ports.OutputBoundary[*models.StatusTag]) (int, *response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("PresentUpdateStatusTag-presenter").End()

	if err := out.Error(); err != nil {
		switch out.Error() {
		case hperror.ErrFieldRequired:
			return http.StatusBadRequest, nil
		default:
			return http.StatusInternalServerError, nil
		}
	}

	resp := out.Response()

	return http.StatusOK, &response.StatusTag{
		ID:     resp.StatusID,
		Status: resp.Status,
	}
}
