package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"

	"encoding/json"

	Config "mainmodule/internal/apps/movies/config"
	"mainmodule/internal/apps/movies/services/movies"
)

func Handler(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	itemID := request.PathParameters["id"]

	// Create item from request
	itemString := request.Body
	itemStruct := movies.Movie{}
	itemStruct.Id = itemID // set the ID

	if err := json.Unmarshal([]byte(itemString), &itemStruct); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	// Get Service and updapte item
	s := movies.MovieService{TableName: Config.Config{}.GetConfig().Movies.TableName}

	item, updateErr := s.CreateOne(&itemStruct)
	if updateErr != nil {
		log.Error().Err(updateErr).Msg("Got error updating item")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	// Marshal item to return
	itemJSON, err := json.Marshal(item)
	if err != nil {
		log.Error().Err(err).Msg("Got error marshalling result")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusCreated, Body: string(itemJSON)}, nil
}

func main() {
	lambda.Start(Handler)
}
