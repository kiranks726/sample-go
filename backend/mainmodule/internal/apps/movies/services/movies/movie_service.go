// MovieService to manage request for Movie data
package movies

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// MovieService struct to contains features
// related to Movie data management
type MovieService struct {
	Name      string
	TableName string
}

func (ps MovieService) DoTest() []string {
	list := []string{"hello", "movies"}
	return list
}

// getDynamoDB will setup session then
// Return the dynamodb client
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

	// TODO: Add 404 for items that are not found

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

// CreateOne to create one Movie,
// Item argument contains the full body of the movie
// Return a status
func (ps MovieService) CreateOne(item *Movie) (Movie, error) {
	// Create DynamoDB client
	svc := ps.getDynamoDB()

	// Set the ID of the newly created item
	itemUUID := uuid.New().String()
	item.Id = itemUUID

	// Marshal to dynamodb item
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Error().Err(err).Msg("DynamoDB MarshalMap Error")
		return *item, err
	}

	// Prepare input for Update Item
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &ps.TableName,
	}

	// UpdateItem request
	_, err = svc.PutItem(input)

	return *item, err
}

// DeleteOne to delete one Movie by ID
// Return a status
func (ps MovieService) DeleteOne(itemID string) error {
	// Create DynamoDB client
	svc := ps.getDynamoDB()

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(itemID),
			},
		},
		TableName: &ps.TableName,
	}

	// DeleteItem request
	_, err := svc.DeleteItem(input)

	return err
}

// UpdateOne to update one Movie by ID, this will do a full overlay of the object
// Item argument contains the ID
// Return a status
func (ps MovieService) UpdateOne(item *Movie) (Movie, error) {
	// Create DynamoDB client
	svc := ps.getDynamoDB()

	// Marshal to dynamodb item
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Error().Err(err).Msg("DynamoDB MarshalMap Error")
		return *item, err
	}

	// Prepare input for Update Item
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &ps.TableName,
	}

	// UpdateItem request
	_, err = svc.PutItem(input)

	return *item, err
}
