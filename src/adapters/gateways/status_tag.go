package gateways

import (
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/params"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
	"gorm.io/gorm"
)

type StatusGateways struct {
	store *gorm.DB
}

func NewStatusGateways(store *gorm.DB) dai.StatusTagRepository {
	return &StatusGateways{
		store: store,
	}
}

func (r *StatusGateways) Create(params.StatusCreate) error {
	return nil
}

func (r *StatusGateways) ReadAll() ([]entities.StatusTag, error) {
	return nil, nil
}
func (r *StatusGateways) Update(params.StatusUpdate) error {
	return nil
}
func (r *StatusGateways) Delete(int) error {
	return nil
}
