package gateways

import (
	"bytes"
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/Hack-Portal/backend/src/usecases/dai"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/redis/go-redis/v9"
)

type cloudflareR2 struct {
	client             *s3.Client
	presignClient      *s3.PresignClient
	bucket             string
	presignLinkExpired time.Duration

	cacheClient dai.Cache[string]
}

// CloudflareR2Option はcloudflareR2のオプションを設定するための関数を定義した型
type CloudflareR2Option func(*cloudflareR2)

// WithPresignLinkExpired はpresignされたリンクの有効期限を設定する
func WithPresignLinkExpired(d time.Duration) CloudflareR2Option {
	return func(c *cloudflareR2) {
		c.presignLinkExpired = d
	}
}

// NewCloudflareR2 はcloudflareR2のインスタンスを生成する
func NewCloudflareR2(bucket string, client *s3.Client, cache *redis.Client, opts ...CloudflareR2Option) dai.FileStore {
	cloudflare := &cloudflareR2{
		bucket:        bucket,
		client:        client,
		cacheClient:   NewCache[string](cache, time.Duration(5)*time.Minute),
		presignClient: s3.NewPresignClient(client),
	}

	for _, opt := range opts {
		opt(cloudflare)
	}

	return cloudflare
}

func checkContentType(file []byte) string {
	return http.DetectContentType(file)
}

func (c *cloudflareR2) ListObjects(ctx context.Context, key string) (*s3.ListObjectsV2Output, error) {
	defer newrelic.FromContext(ctx).StartSegment("ListObjects-gateway").End()

	return c.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(c.bucket),
		Prefix: aws.String(key),
	})
}

func (c *cloudflareR2) GetPresignedObjectURL(ctx context.Context, key string) (string, error) {
	defer newrelic.FromContext(ctx).StartSegment("GetPresignedObjectURL-gateway").End()
	url, err := c.cacheClient.Get(ctx, key, func(ctx context.Context) (string, error) {
		object, err := c.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
			Bucket:          aws.String(c.bucket),
			Key:             aws.String(key),
			ResponseExpires: aws.Time(time.Now().Add(c.presignLinkExpired * time.Hour)),
		})
		if err != nil {
			return "", err
		}
		return object.URL, nil
	})

	return url, err
}

func (c *cloudflareR2) UploadFile(ctx context.Context, file []byte, key string) (string, error) {
	defer newrelic.FromContext(ctx).StartSegment("UploadFile-gateway").End()

	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(c.bucket),
		Key:           aws.String(key),
		Body:          bytes.NewReader(file),
		ContentLength: aws.Int64(int64(len(file))),
		ContentType:   aws.String(checkContentType(file)),
	})

	return key, err
}

func (c *cloudflareR2) DeleteFile(ctx context.Context, fileName string) error {
	defer newrelic.FromContext(ctx).StartSegment("DeleteFile-gateway").End()

	_, err := c.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(fileName),
	})
	return err
}

func (c *cloudflareR2) ParallelGetPresignedObjectURL(ctx context.Context, input []dai.ParallelGetPresignedObjectURLInput) (map[string]string, error) {
	defer newrelic.FromContext(ctx).StartSegment("ParallelGetPresignedObjectURL-gateway").End()

	type resultCh struct {
		hackathonID string
		url         string
	}

	var (
		output = make(map[string]string, len(input))
		ch     = make(chan resultCh, len(input))
		wg     sync.WaitGroup
	)
	defer close(ch)

	for _, in := range input {
		wg.Add(1)
		go func(in dai.ParallelGetPresignedObjectURLInput) {
			defer wg.Done()
			url, err := c.GetPresignedObjectURL(ctx, in.Key)
			if err != nil {
				ch <- resultCh{
					hackathonID: in.HackathonID,
					url:         "",
				}
			} else {
				ch <- resultCh{
					hackathonID: in.HackathonID,
					url:         url,
				}
			}
		}(in)
	}

	wg.Wait()

	for range input {
		result := <-ch
		output[result.hackathonID] = result.url
	}

	return output, nil
}
