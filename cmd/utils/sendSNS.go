package snsmananger

import (
	"sms-aws-go/cmd/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendSMS(Message string, PhoneNumber string) error {
	config, err := config.GetAwsConfig()

	if err != nil {
		return err
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.AccessKeyID, config.SecretAccessKey, ""),
	})
	if err != nil {
		return err
	}

	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{
		Message:     aws.String(Message),
		PhoneNumber: aws.String(PhoneNumber),
	})
	if err != nil {
		return err
	}

	_ = result // Apenas para suprimir o aviso do compilador

	return nil
}
