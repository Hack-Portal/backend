package params

type StatusCreate struct {
	Language string
	Icon     string
}
type StatusUpdate struct {
	StatusTagID int32
	Language    string
	Icon        string
}
