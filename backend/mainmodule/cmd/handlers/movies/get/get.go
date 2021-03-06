package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"encoding/json"

	Config "mainmodule/internal/apps/movies/config"
	"mainmodule/internal/apps/movies/services/movies"
)

func Handler(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Getting id from path parameters
	pathParamID := request.PathParameters["id"]

	// Load app configuration
	config := Config.Config{}.GetConfig()

	// GetItem request
	s := movies.MovieService{TableName: config.Stacks.Movies.Tablename}
	movie, err := s.FindOne(pathParamID)

	// Checking for errors, return error
	if err != nil {
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	// Marshal to type []uint8
	marshalledItem, err := json.Marshal(movie)
	if err != nil {
		panic(fmt.Sprint("Failed to UnmarshalMap result.Item: ", err))
	}

	// TODO: Make standard response object for REST
	// Return marshalled item
	response := events.APIGatewayProxyResponse{Body: string(marshalledItem), StatusCode: http.StatusOK}
	response.Headers = make(map[string]string)
	response.Headers["Content-Type"] = "application/json"

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
