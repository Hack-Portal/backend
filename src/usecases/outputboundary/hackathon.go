package outputboundary

// ここでは主にPresenterを汎化している

// TODO:ここもどうようにするか考える
type HackathonOutputPort interface {
	Create(error)
	Read()
	Update()
	Delete()
}
