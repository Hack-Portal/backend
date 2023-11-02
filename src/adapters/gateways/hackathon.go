package gateways

import (
	"temp/src/datastructs/entities"
	"temp/src/datastructs/params"
	"temp/src/usecases/dai"

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

func (h *HackathonGateway) Create(args params.HackathonCreate) error {
	return h.store.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&args.Hackathon).Error; err != nil {
			return err
		}

		var hackathonStatusTags []entities.HackathonStatusTag
		for _, tag := range args.Statuses {
			hackathonStatusTags = append(hackathonStatusTags, entities.HackathonStatusTag{
				HackathonID: args.Hackathon.HackathonID,
				StatusID:    tag,
			})
		}

		if err := tx.Create(&hackathonStatusTags).Error; err != nil {
			return err
		}

		return nil
	})
}

func (h *HackathonGateway) Read()   {}
func (h *HackathonGateway) Update() {}
func (h *HackathonGateway) Delete() {}
