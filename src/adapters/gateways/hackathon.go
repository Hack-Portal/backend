package gateways

import (
	"time"

	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/params"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
	"gorm.io/gorm"
)

// ここでは、daiで定義したinterfaceを実装する

type HackathonGateway struct {
	store *gorm.DB
}

func NewHackathonGateway(store *gorm.DB) dai.HackathonRepository {
	return &HackathonGateway{
		store: store,
	}
}

func (h *HackathonGateway) Create(arg params.HackathonCreate) error {
	return h.store.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&arg.Hackathon).Error; err != nil {
			return err
		}

		var hackathonStatusTags []entities.HackathonStatusTag
		for _, tag := range arg.Statuses {
			hackathonStatusTags = append(hackathonStatusTags, entities.HackathonStatusTag{
				HackathonID: arg.Hackathon.HackathonID,
				StatusID:    tag,
			})
		}

		if err := tx.Create(&hackathonStatusTags).Error; err != nil {
			return err
		}
		return nil
	})
}

func (h *HackathonGateway) ReadAll(arg params.HackathonReadAll) ([]entities.Hackathon, []entities.HackathonStatus, error) {
	var (
		hackathons []entities.Hackathon
		statusTags []entities.HackathonStatus
		err        error
	)

	if len(arg.SortTag) == 0 {
		err = h.store.Where("hackathons.expired > ? AND hackathons.is_delete = ?", time.Now(), false).
			Limit(arg.Limit).
			Offset(arg.Offset).
			Find(&hackathons).
			Error
	} else {
		err = h.store.Joins("JOIN hackathon_status_tags ON hackathons.hackathon_id = hackathon_status_tags.hackathon_id AND hackathon_status_tags.status_id IN ?", arg.SortTag).
			Where("hackathons.expired > ? AND hackathons.is_delete = ?", time.Now(), false).
			Limit(arg.Limit).
			Offset(arg.Offset).
			Find(&hackathons).
			Error
	}
	if err != nil {
		return nil, nil, err
	}

	var hackathonIDs []string

	for _, hackathon := range hackathons {
		hackathonIDs = append(hackathonIDs, hackathon.HackathonID)
	}

	rows, err := h.store.Joins("JOIN status_tags ON hackathon_status_tags.status_id = status_tags.status_id").
		Where("hackathon_status_tags.hackathon_id IN ?", hackathonIDs).
		Select("hackathon_status_tags.hackathon_id,status_tags.status_id,status_tags.status").
		Rows()
	if err != nil {
		return nil, nil, err
	}

	for rows.Next() {
		var statusTag entities.HackathonStatus
		if err := rows.Scan(&statusTag.HackathonID, &statusTag.StatusID, &statusTag.Status); err != nil {
			return nil, nil, err
		}
		statusTags = append(statusTags, statusTag)
	}

	return hackathons, statusTags, nil
}
func (h *HackathonGateway) Update() {}
func (h *HackathonGateway) Delete() {}
