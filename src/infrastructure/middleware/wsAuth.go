package middleware

import (
	"encoding/base64"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type WsAuthRequest struct {
	ProjectID string `json:"project_id"`
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	AuthToken string `json:"auth_token"`
}

const (
	WS_REQUEST = "ws_request"
)

func (m *middleware) WsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		wsauth, err := m.decodeURI(ctx.Param("token"))
		if err != nil {
			m.l.Infof("invalid request : %v", err)
			ctx.Abort()
			return
		}
		ctx.Set(WS_REQUEST, wsauth)
		ctx.Next()
	}
}

// // URIデコードを一旦関数化
func (m *middleware) decodeURI(uri string) (wsauth *WsAuthRequest, err error) {
	decodedUri, err := base64.URLEncoding.DecodeString(uri)
	if err != nil {
		m.l.Infof("itnvalid request : %v", err)
		return
	}
	if err = json.Unmarshal(decodedUri, &wsauth); err != nil {
		m.l.Infof("invalid request : %v", err)
		return
	}
	return
}
