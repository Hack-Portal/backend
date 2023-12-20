package gateways

import (
	"bytes"
	"context"
	"net/http"

	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type CloudflareR2 struct {
	client        *s3.Client
	PresignClient *s3.PresignClient
	bucket        string
}

func NewCloudflareR2(bucket string, client *s3.Client) dai.FileStore {
	return &CloudflareR2{
		bucket:        bucket,
		client:        client,
		PresignClient: s3.NewPresignClient(client),
	}
}

func checkContentType(file []byte) string {
	return http.DetectContentType(file)
}

func (c *CloudflareR2) UploadFile(ctx context.Context, file []byte, key string) (string, error) {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(c.bucket),
		Key:           aws.String(key),
		Body:          bytes.NewReader(file),
		ContentLength: aws.Int64(int64(len(file))),
		ContentType:   aws.String(checkContentType(file)),
	})
	if err != nil {
		return "", err
	}

	object, err := c.PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}

	return object.URL, nil
}

func (c *CloudflareR2) DeleteFile(ctx context.Context, fileName string) error {
	_, err := c.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(fileName),
	})
	return err
}
