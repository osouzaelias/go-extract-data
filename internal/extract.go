package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-extract-data/pkg/aws"
	"go-extract-data/pkg/config"
)

type ExtractData struct {
	storage  *aws.S3Client
	database *aws.DynamoDBClient
	config   config.Config
}

func NewExtractData(c config.Config) *ExtractData {
	return &ExtractData{
		storage:  aws.NewS3Client(c),
		database: aws.NewDynamoDBClient(c),
		config:   c,
	}
}

func (d ExtractData) Extract(bucketTypeName config.BucketTypeName) error {
	results, err := d.database.ExecuteStatement(d.config.Query(bucketTypeName))

	if err != nil {
		return fmt.Errorf("falha na query do DynamoDB: %s", err)
	}

	for index, result := range results {
		jsonData, _ := json.Marshal(result)

		reader := bytes.NewReader(jsonData)

		err = d.storage.PutObject(d.config.BucketName(bucketTypeName), fmt.Sprintf("item_%d.json", index), reader)

		if err != nil {
			return fmt.Errorf("falha ao enviar objeto para S3: %s", err)
		}
	}

	return nil
}
