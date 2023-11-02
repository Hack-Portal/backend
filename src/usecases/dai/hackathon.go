package dai

import (
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/entities"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/params"
)

// 正式名 Data Access Interfaceと呼ばれ、Gatewaysに定義するデータの永続化などについて汎化する

// TODO:ここもどのようにするか考える
type HackathonRepository interface {
	Create(params.HackathonCreate) error
	ReadAll(params.HackathonReadAll) ([]entities.Hackathon, []entities.HackathonStatus, error)
	Update()
	Delete()
}
