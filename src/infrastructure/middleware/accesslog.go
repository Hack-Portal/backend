package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

type AccessLog struct {
	Request      string `json:"request"`
	Response     string `json:"response"`
	Status       int    `json:"status"`
	ContentLengh int64  `json:"content_length"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	Query        string `json:"query"`
	IP           string `json:"ip"`
	UserAgent    string `json:"user_agent"`
	Latency      string `json:"latency"`
}

func (m *middleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		var buf bytes.Buffer
		body, _ := io.ReadAll(io.TeeReader(c.Request.Body, &buf))
		c.Request.Body = io.NopCloser(&buf)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		body, err := json.Marshal(AccessLog{
			Request:      string(body),
			Response:     blw.body.String(),
			Status:       c.Writer.Status(),
			ContentLengh: c.Request.ContentLength,
			Method:       c.Request.Method,
			Path:         c.Request.URL.Path,
			Query:        c.Request.URL.RawQuery,
			IP:           c.ClientIP(),
			UserAgent:    c.Request.UserAgent(),
			Latency:      time.Since(start).String(),
		})
		if err != nil {
			c.AbortWithStatusJSON(500, nil)
			m.l.Error(err)
			return
		}

		m.l.Infof(`request mes="%s"`, string(body))
	}
}
