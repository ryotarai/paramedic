package store

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ryotarai/paramedic/awsclient"
)

const commandsTableName = "ParamedicCommands"

type Store struct {
	dynamodb awsclient.DynamoDB
}

func New(dynamodb awsclient.DynamoDB) *Store {
	return &Store{
		dynamodb: dynamodb,
	}
}

func (s *Store) PutCommand(r *CommandRecord) error {
	av, err := dynamodbattribute.MarshalMap(*r)
	if err != nil {
		return err
	}

	_, err = s.dynamodb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(commandsTableName),
		Item:      av,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetCommand(commandID string) (*CommandRecord, error) {
	resp, err := s.dynamodb.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(commandsTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"CommandID": {S: aws.String(commandID)},
		},
	})
	if err != nil {
		return nil, err
	}

	r := CommandRecord{}
	err = dynamodbattribute.UnmarshalMap(resp.Item, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (s *Store) CreateTablesIfNotExists() error {
	found := false
	err := s.dynamodb.ListTablesPages(&dynamodb.ListTablesInput{}, func(resp *dynamodb.ListTablesOutput, last bool) bool {
		for _, n := range resp.TableNames {
			if *n == commandsTableName {
				found = true
				return false
			}
		}
		return true
	})
	if err != nil {
		return err
	}

	if found {
		return nil
	}

	return s.createTables()
}

func (s *Store) createTables() error {
	log.Printf("INFO: creating %s table", commandsTableName)
	_, err := s.dynamodb.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(commandsTableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("CommandID"),
				AttributeType: aws.String("S"), // string
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("CommandID"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	})
	return err
}
