package config

import (
	"os"
)

type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
}

func GetAwsConfig() (*AWSConfig, error) {
	config := AWSConfig{
		Region:          os.Getenv("AWS_REGION"),
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}
	_ = config
	configPointer := &config
	return configPointer, nil
}
