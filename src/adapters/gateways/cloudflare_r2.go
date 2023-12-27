package gateways

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	hackathonPrefix = "hackathon"
)

type CloudflareR2 struct {
	client        *s3.Client
	PresignClient *s3.PresignClient
	bucket        string
	Config        Config
}

type Config struct {
	PresignLinkExpired time.Duration
}

func NewCloudflareR2(bucket string, client *s3.Client, presignLinkExpired int) dai.FileStore {
	return &CloudflareR2{
		bucket:        bucket,
		client:        client,
		PresignClient: s3.NewPresignClient(client),
		Config: Config{
			PresignLinkExpired: time.Duration(presignLinkExpired),
		},
	}
}

func checkContentType(file []byte) string {
	return http.DetectContentType(file)
}

func (c *CloudflareR2) getObject(ctx context.Context, key string) (*s3.GetObjectOutput, error) {
	return c.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
}

func (c *CloudflareR2) ListObjects(ctx context.Context, key string) (*s3.ListObjectsV2Output, error) {
	return c.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(c.bucket),
		Prefix: aws.String(key),
	})
}

func (c *CloudflareR2) GetPresignedObjectURL(ctx context.Context, key string) (string, error) {
	object, err := c.PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket:          aws.String(c.bucket),
		Key:             aws.String(key),
		ResponseExpires: aws.Time(time.Now().Add(c.Config.PresignLinkExpired * time.Hour)),
	})
	if err != nil {
		return "", err
	}

	return object.URL, nil
}

func (c *CloudflareR2) UploadFile(ctx context.Context, file []byte, key string) (string, error) {
	objectKey := fmt.Sprintf("%s/%s", hackathonPrefix, key)
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(c.bucket),
		Key:           aws.String(key),
		Body:          bytes.NewReader(file),
		ContentLength: aws.Int64(int64(len(file))),
		ContentType:   aws.String(checkContentType(file)),
	})

	return objectKey, err
}

func (c *CloudflareR2) DeleteFile(ctx context.Context, fileName string) error {
	_, err := c.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(fileName),
	})
	return err
}
