package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"go-extract-data/pkg/config"
	"log"
)

type DynamoDBClient struct {
	client *dynamodb.Client
}

func NewDynamoDBClient(c config.Config) *DynamoDBClient {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(c.Region))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}
	return &DynamoDBClient{
		client: dynamodb.NewFromConfig(cfg),
	}
}

func (d *DynamoDBClient) ExecuteStatement(statement string) ([]map[string]interface{}, error) {
	params := &dynamodb.ExecuteStatementInput{
		Statement: aws.String(statement),
	}
	resp, err := d.client.ExecuteStatement(context.TODO(), params)
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	err = attributevalue.UnmarshalListOfMaps(resp.Items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
