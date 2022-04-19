package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

	"encoding/json"
	"fmt"

	Config "github.com/kiranks726/sample-go/src/apps/movies/config"
	"github.com/kiranks726/sample-go/src/apps/movies/services/movies"
)

func Handler(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config := Config.Config{}.GetConfig()
	// Creating session for client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// New uuid for item id
	itemUUID := uuid.New().String()

	fmt.Println("Generated new item uuid:", itemUUID)

	// Create item from request
	itemString := request.Body
	itemStruct := movies.Movie{}
	itemStruct.Id = itemUUID

	if err := json.Unmarshal([]byte(itemString), &itemStruct); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	// Validate, item must have a title
	if itemStruct.Title == "" {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	fmt.Printf("after UUID: %v", itemStruct)

	// Marshal to dynamodb item
	av, err := dynamodbattribute.MarshalMap(itemStruct)
	if err != nil {
		fmt.Println("Error marshalling item: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	tableName := config.Movies.TableName
	fmt.Printf("tableName: %v", tableName)

	// Build put item input
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	fmt.Println("Create item from input: ")
	fmt.Println(input)

	// PutItem request
	_, err = svc.PutItem(input)

	// Checking for errors, return error
	if err != nil {
		fmt.Println("Got error calling PutItem: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	// Marshal item to return
	itemJSON, err := json.Marshal(itemStruct)
	if err != nil {
		fmt.Println("Got error marshalling result: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	fmt.Println("Returning item: ", string(itemJSON))

	// Returning response with AWS Lambda Proxy Response
	response := events.APIGatewayProxyResponse{Body: string(itemJSON), StatusCode: http.StatusCreated}
	response.Headers = make(map[string]string)
	response.Headers["Content-Type"] = "application/json"

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
