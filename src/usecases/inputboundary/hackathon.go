package inputboundary

import "github.com/hackhack-Geek-vol6/backend/src/datastructs/input"

// ここでは主にUsecase Interactorを汎化している

// TODO:しっかりとしたインターフェースを作る　一旦は仮置き
type HackathonInputPort interface {
	Create(input.HackathonCreate, []byte)
	Read()
	Update()
	Delete()
}