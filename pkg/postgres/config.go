package postgres

type Status int

const (
	UNKNOWN Status = iota
	NOT_READY
	READY
	DISCONNECTED
	ERROR
)
