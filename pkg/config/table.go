package config

import "os"

type Table struct {
	dynamoDb string
	mySql    string
}

func NewTable() *Table {
	return &Table{
		dynamoDb: os.Getenv("TABLE_DYNAMODB"),
		mySql:    os.Getenv("TABLE_MYSQL"),
	}
}

func (t Table) DynamoDb() string {
	return t.dynamoDb
}

func (t Table) MySql() string {
	return t.mySql
}
