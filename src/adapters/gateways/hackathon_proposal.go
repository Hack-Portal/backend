package gateways

import (
	"context"
	"time"

	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type HackathonProposalGateway struct {
	db          *gorm.DB
	cacheClient dai.Cache[[]*models.HackathonProposal]
}

func NewHackathonProposalGateway(db *gorm.DB, cache *redis.Client) dai.HackathonProposalDai {
	return &HackathonProposalGateway{
		db:          db,
		cacheClient: NewCache[[]*models.HackathonProposal](cache, time.Duration(5)*time.Minute),
	}
}

func (hpg *HackathonProposalGateway) Create(ctx context.Context, url string) (id int64, err error) {
	defer newrelic.FromContext(ctx).StartSegment("CreateHackathonProposal-gateway").End()

	result := hpg.db.Create(&models.HackathonProposal{
		URL: url,
	})
	if result.Error != nil {
		return 0, result.Error
	}

	var hackathonProposalID int64
	err = hpg.db.Raw("SELECT currval(pg_get_serial_sequence('hackathon_proposals', 'hackathon_proposal_id'))").Scan(&hackathonProposalID).Error
	if err != nil {
		return 0, err
	}

	return hackathonProposalID, hpg.cacheClient.Reset(ctx, "hackathon_proposals")
}

// func (h *HackathonProposalGateway) FindAll(ctx context.Context, arg request.ListHackathonProposal) ([]*models.HackathonProposal, error) {
// 	defer newrelic.FromContext(ctx).StartSegment("FindAllHackathonProposal-gateway").End()

// 	var key string = "hackathon-proposal"
// 	chain := h.db.Limit(arg.PageSize).Offset(arg.PageID)

// 	// if len(arg.)
// 	hackathonProposals, err := h.cacheClient.Get(ctx, key, func(ctx context.Context) ([]*models.HackathonProposal, error) {
// 		var hackathonProposals []*models.HackathonProposal
// 		result := chain.Select("DISTINCT (hackathon)", "url", "is_approval", "created_at").Find(&hackathonProposals)
// 		if result.Error != nil {
// 			return nil, result.Error
// 		}

// 		return hackathonProposals, nil
// 	})
// 	return hackathonProposals, err
// }

// func (h *HackathonProposalGateway) Delete(ctx context.Context, hackathonProposalID string) error {
// 	defer newrelic.FromContext(ctx).StartSegment("DeleteHackathonProposal-gateway").End()

// 	result := h.db.Delete(&models.HackathonProposal{}, "hackathon_proposal_id = ?", hackathonProposalID)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return h.cacheClient.Reset(ctx, "hackathon_proposals")
// }