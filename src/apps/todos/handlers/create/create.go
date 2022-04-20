package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Id      string `json:"Id,omitempty"`
	Title   string `json:"Title"`
	Details string `json:"Details"`
	Message string `json:"Message"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Creating session for client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// New uuid for item id
	itemUuid := uuid.New().String()

	fmt.Println("Generated new item uuid:", itemUuid)

	// Unmarshal to Item to access object properties
	itemString := request.Body
	itemStruct := Item{}
	json.Unmarshal([]byte(itemString), &itemStruct)

	if itemStruct.Title == "" {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// Create new item of type item
	var item = new(Item)
	item.Id = itemUuid
	item.Title = itemStruct.Title
	item.Details = itemStruct.Details

	// Marshal to dynamodb item
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Error marshalling item: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	tableName := os.Getenv("todoTableName")

	// Build put item input
	fmt.Println("Putting item: %+v", av)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	// PutItem request
	_, err = svc.PutItem(input)

	// Checking for errors, return error
	if err != nil {
		fmt.Println("Got error calling PutItem: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	// Marshal item to return
	itemMarshalled, err := json.Marshal(item)

	fmt.Println("Returning item: ", string(itemMarshalled))

	//Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{Body: string(itemMarshalled), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
