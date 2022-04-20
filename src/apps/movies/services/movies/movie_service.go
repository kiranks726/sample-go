// MovieService to manage request for Movie data
package movies

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/zerolog/log"
)

// MovieService struct to contains features
// related to Movie data management
type MovieService struct {
	Name      string
	TableName string
}

func (ps MovieService) getDynamoDB() *dynamodb.DynamoDB {
	// Creating session for client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return svc
}

// FindList to get a list of Movies
// Return an array of Movies
func (ps MovieService) FindList() ([]Movie, error) {
	// Create DynamoDB client
	svc := ps.getDynamoDB()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: &ps.TableName,
	}

	// Scan table
	result, err := svc.Scan(params)

	// Checking for errors, return error
	if err != nil {
		log.Error().Err(err).Msg("Query API call failed")
		return []Movie{}, err
	}

	// Setup result data
	var itemArray []Movie

	for _, i := range result.Items {
		item := Movie{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Error().Err(err).Msg("DynamoDB UnmarshalError")
		}

		itemArray = append(itemArray, item)
	}

	return itemArray, err
}

// FindOne to get a single Movie by ID
// Return a single Movie
func (ps MovieService) FindOne(itemID string) (Movie, error) {
	// Create DynamoDB client
	svc := ps.getDynamoDB()

	// GetItem request
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: &ps.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(itemID),
			},
		},
	})

	// Checking for errors, return error
	if err != nil {
		log.Error().Err(err).Msg("DynamoDB GetItem Error")
		return Movie{}, err
	}

	// Created item of type Item
	item := Movie{}

	// UnmarshallMap result.item into item
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		log.Error().Err(err).Msg("Failed to UnmarshalMap result.Item")
	}

	return item, err
}
