package output

// ここでは、返す構造体を定義する
// 所謂レスポンスボディ等の構造体を定義する

type CreateHackathon struct {
	Message string
}

type HackathonRead struct{}
type HackathonUpdate struct{}
type HackathonDelete struct{}
