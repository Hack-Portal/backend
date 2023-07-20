package api

import (
	"github.com/gin-gonic/gin"
)

const (
	MultiPartNextPartEoF   = "multipart: NextPart: EOF"
	RequestContentTypeIsnt = "request Content-Type isn't multipart/form-data"
	HttpNoSuchFile         = "http: no such file"
	ImageKey               = "icons"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
