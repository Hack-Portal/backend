package dai

import (
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/params"
)

type StatusTagRepository interface {
	Create(params.StatusCreate) error
	ReadAll() ([]entities.StatusTag, error)
	Update(params.StatusUpdate) error
	Delete(int) error
}
