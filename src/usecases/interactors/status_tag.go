package interactors

import (
	"context"

	"github.com/Hack-Portal/backend/src/datastructure/hperror"
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type statusTagInteractor struct {
	StatusTagRepo dai.StatusTagDai
	discordNotify DiscordNotify
	Output        ports.StatusTagOutputBoundary
}

// NewStatusTagInteractor はStatusTagに関するユースケースを生成します
func NewStatusTagInteractor(statusTagRepo dai.StatusTagDai, discordNotify DiscordNotify, output ports.StatusTagOutputBoundary) ports.StatusTagInputBoundary {
	return &statusTagInteractor{
		StatusTagRepo: statusTagRepo,
		discordNotify: discordNotify,
		Output:        output,
	}
}

// CreateStatusTag はStatusTagを作成します
func (s *statusTagInteractor) CreateStatusTag(ctx context.Context, in *request.CreateStatusTag) (int, *response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("CreateStatusTag-usecase").End()

	if in.Status == "" {
		return s.Output.PresentCreateStatusTag(ctx, ports.NewOutput[*models.StatusTag](
			hperror.ErrFieldRequired,
			nil,
		))
	}

	id, err := s.StatusTagRepo.Create(ctx, &models.StatusTag{
		Status: in.Status,
	})
	if err != nil {
		return s.Output.PresentCreateStatusTag(ctx, ports.NewOutput[*models.StatusTag](
			err,
			nil,
		))
	}

	if err := s.discordNotify.CreateNewForumTag([]*models.StatusTag{
		{
			StatusID: id,
			Status:   in.Status,
		},
	}); err != nil {
		return s.Output.PresentCreateStatusTag(ctx, ports.NewOutput[*models.StatusTag](
			err,
			nil,
		))
	}

	return s.Output.PresentCreateStatusTag(ctx, ports.NewOutput[*models.StatusTag](
		nil,
		&models.StatusTag{
			StatusID: id,
			Status:   in.Status,
		},
	))
}

// FindAllStatusTag はStatusTagを全て取得します
func (s *statusTagInteractor) FindAllStatusTag(ctx context.Context) (int, []*response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("FindAllStatusTag-usecase").End()

	statusTags, err := s.StatusTagRepo.FindAll(ctx)
	return s.Output.PresentFindAllStatusTag(ctx, ports.NewOutput[[]*models.StatusTag](
		err,
		statusTags,
	))
}

// FindByIDStatusTag はStatusTagをIDで取得します
func (s *statusTagInteractor) FindByIDStatusTag(ctx context.Context, in *request.GetStatusTagByID) (int, *response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("FindByIdStatusTag-usecase").End()

	statusTag, err := s.StatusTagRepo.FindByID(ctx, in.ID)
	return s.Output.PresentFindByIDStatusTag(ctx, ports.NewOutput[*models.StatusTag](
		err,
		statusTag,
	))
}

// UpdateStatusTag はStatusTagを更新します
func (s *statusTagInteractor) UpdateStatusTag(ctx context.Context, in *request.UpdateStatusTag) (int, *response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("UpdateStatusTag-usecase").End()

	if in.Status == "" {
		return s.Output.PresentUpdateStatusTag(ctx, ports.NewOutput[*models.StatusTag](
			hperror.ErrFieldInvalid,
			nil,
		))
	}

	id, err := s.StatusTagRepo.Update(ctx, &models.StatusTag{
		StatusID: in.ID,
		Status:   in.Status,
	})
	return s.Output.PresentUpdateStatusTag(ctx, ports.NewOutput[*models.StatusTag](
		err,
		&models.StatusTag{
			StatusID: id,
			Status:   in.Status,
		},
	))
}
