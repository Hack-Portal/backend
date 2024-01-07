package v1

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type v1router struct {
	v1 *echo.Group

	db     *gorm.DB
	cache  *redis.Client
	client *s3.Client
}

func NewV1Router(e *echo.Group, db *gorm.DB, cache *redis.Client, client *s3.Client) {
	router := &v1router{
		v1: e,

		db:     db,
		cache:  cache,
		client: client,
	}

	router.statusTag()
	router.hackathon()
	return
}
