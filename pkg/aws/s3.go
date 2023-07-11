package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go-extract-data/pkg/config"
	"io"
	"log"
)

type S3Client struct {
	client *s3.Client
}

func NewS3Client(c config.Config) *S3Client {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(c.Region))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}
	return &S3Client{
		client: s3.NewFromConfig(cfg),
	}
}

func (s *S3Client) PutObject(bucketName, key string, body io.Reader) error {
	putParams := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   body,
	}
	_, err := s.client.PutObject(context.TODO(), putParams)
	if err != nil {
		return err
	}
	return nil
}
