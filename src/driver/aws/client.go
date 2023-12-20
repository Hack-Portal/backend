package aws

import (
	"context"
	"fmt"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type awsConnect struct {
}

type AwsConnect interface {
	Connect(ctx context.Context) (*s3.Client, error)
}

func New() AwsConnect {
	return &awsConnect{}
}

func (a *awsConnect) newResolver() aws.EndpointResolverWithOptions {
	return aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               fmt.Sprintf("https://%s.%s", config.Config.Buckets.AccountID, config.Config.Buckets.EndPoint),
			HostnameImmutable: true,
			Source:            aws.EndpointSourceCustom,
		}, nil
	})
}

func (a *awsConnect) Connect(ctx context.Context) (*s3.Client, error) {
	cfg, err := awsConfig.LoadDefaultConfig(
		ctx,
		awsConfig.WithEndpointResolverWithOptions(a.newResolver()),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			config.Config.Buckets.AccessKeyId,
			config.Config.Buckets.AccessKeySecret,
			"",
		)),
		awsConfig.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}
