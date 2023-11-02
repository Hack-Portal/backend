package dai

import "temp/src/datastructs/params"

// 正式名 Data Access Interfaceと呼ばれ、Gatewaysに定義するデータの永続化などについて汎化する

// TODO:ここもどのようにするか考える
type HackathonRepository interface {
	Create(params.HackathonCreate) error
	Read()
	Update()
	Delete()
}
