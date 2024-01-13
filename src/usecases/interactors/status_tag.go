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

type StatusTagInteractor struct {
	StatusTagRepo dai.StatusTagDai
	discordNotify DiscordNotify
	Output        ports.StatusTagOutputBoundary
}

func NewStatusTagInteractor(statusTagRepo dai.StatusTagDai, discordNotify DiscordNotify, output ports.StatusTagOutputBoundary) ports.StatusTagInputBoundary {
	return &StatusTagInteractor{
		StatusTagRepo: statusTagRepo,
		discordNotify: discordNotify,
		Output:        output,
	}
}

func (s *StatusTagInteractor) CreateStatusTag(ctx context.Context, in *request.CreateStatusTag) (int, *response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("CreateStatusTag-usecase").End()

	if in.Status == "" {
		return s.Output.PresentCreateStatusTag(ctx, &ports.OutputCraeteStatusTagData{
			Error:    hperror.ErrFieldRequired,
			Response: nil,
		})
	}

	id, err := s.StatusTagRepo.Create(ctx, &models.StatusTag{
		Status: in.Status,
	})

	if err := s.discordNotify.CreateNewForumTag([]*models.StatusTag{
		{
			StatusID: id,
			Status:   in.Status,
		},
	}); err != nil {
		return s.Output.PresentCreateStatusTag(ctx, &ports.OutputCraeteStatusTagData{
			Error:    err,
			Response: nil,
		})
	}

	return s.Output.PresentCreateStatusTag(ctx, &ports.OutputCraeteStatusTagData{
		Error: err,
		Response: &models.StatusTag{
			StatusID: id,
			Status:   in.Status,
		},
	})
}

func (s *StatusTagInteractor) FindAllStatusTag(ctx context.Context) (int, []*response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("FindAllStatusTag-usecase").End()

	statusTags, err := s.StatusTagRepo.FindAll(ctx)
	return s.Output.PresentFindAllStatusTag(ctx, &ports.OutputFindAllStatusTagData{
		Error:    err,
		Response: statusTags,
	})
}

func (s *StatusTagInteractor) FindByIdStatusTag(ctx context.Context, in *request.GetStatusTagByID) (int, *response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("FindByIdStatusTag-usecase").End()

	statusTag, err := s.StatusTagRepo.FindById(ctx, in.ID)
	return s.Output.PresentFindByIdStatusTag(ctx, &ports.OutputFindByIdStatusTagData{
		Error:    err,
		Response: statusTag,
	})
}

func (s *StatusTagInteractor) UpdateStatusTag(ctx context.Context, in *request.UpdateStatusTag) (int, *response.StatusTag) {
	defer newrelic.FromContext(ctx).StartSegment("UpdateStatusTag-usecase").End()

	if in.Status == "" {
		return s.Output.PresentUpdateStatusTag(ctx, &ports.OutputUpdateStatusTagData{
			Error:    hperror.ErrFieldRequired,
			Response: nil,
		})
	}

	id, err := s.StatusTagRepo.Update(ctx, &models.StatusTag{
		StatusID: in.ID,
		Status:   in.Status,
	})
	return s.Output.PresentUpdateStatusTag(ctx, &ports.OutputUpdateStatusTagData{
		Error: err,
		Response: &models.StatusTag{
			StatusID: id,
			Status:   in.Status,
		},
	})
}
