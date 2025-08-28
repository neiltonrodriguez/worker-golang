package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func LoadAWSConfig() (cfg aws.Config, err error) {
	cfg, err = config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(GlobalConfig.AWS.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			GlobalConfig.AWS.AccessKey,
			GlobalConfig.AWS.SecretKey,
			"",
		)),
	)
	return
}
